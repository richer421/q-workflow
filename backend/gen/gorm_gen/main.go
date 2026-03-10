package main

import (
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./infra/mysql/dao",
		Mode:    gen.WithDefaultQuery,
	})

	// TODO: Add business models
	// g.ApplyBasic(model.Workflow{})

	g.Execute()
}
