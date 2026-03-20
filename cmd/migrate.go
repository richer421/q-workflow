package cmd

import (
	"github.com/richer421/q-workflow/conf"
	"github.com/richer421/q-workflow/infra/mysql"
	"github.com/richer421/q-workflow/pkg/logger"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		defer mysql.Close()

		logger.Infof("Running migration...")
		if err := mysql.Migrate(conf.C.MySQL); err != nil {
			logger.Fatalf("migration failed: %v", err)
		}
		logger.Infof("Migration completed!")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
