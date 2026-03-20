package mysql

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/richer421/q-workflow/conf"
	"github.com/richer421/q-workflow/infra/mysql/dao"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dbNameRe = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

func Init(cfg conf.MySQLConfig) error {
	if DB != nil {
		return nil
	}

	if cfg.Database == "" {
		return fmt.Errorf("database config is required")
	}

	// 连接到指定数据库
	db, err := gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		return err
	}

	DB = db
	dao.SetDefault(db)
	return nil
}

func Migrate(cfg conf.MySQLConfig) error {
	if cfg.Database == "" {
		return fmt.Errorf("database config is required")
	}

	createDB, err := gorm.Open(mysql.Open(cfg.DSNWithoutDB()), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlCreateDB, err := createDB.DB()
	if err != nil {
		return err
	}
	defer sqlCreateDB.Close()

	if err := ensureDatabase(createDB, cfg.Database); err != nil {
		return err
	}

	if err := Init(cfg); err != nil {
		return err
	}

	// TODO: Add business models
	return nil
}

func ensureDatabase(db *gorm.DB, name string) error {
	if !dbNameRe.MatchString(name) {
		return fmt.Errorf("invalid database name: %s", name)
	}

	sql := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci",
		strings.ReplaceAll(name, "`", "``"),
	)
	return db.Exec(sql).Error
}

func Close() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
