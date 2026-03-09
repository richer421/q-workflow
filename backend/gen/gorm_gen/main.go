package main

import (
	"github.com/richer/q-workflow/infra/mysql/model"

	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./infra/mysql/dao",
		Mode:    gen.WithDefaultQuery,
	})

	g.ApplyBasic(model.HelloWorld{})

	g.Execute()
}
