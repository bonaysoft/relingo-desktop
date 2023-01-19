package main

import (
	"github.com/bonaysoft/relingo-desktop/pkg/dal/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./pkg/dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(sqlite.Open("gorm.db"))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(model.Word{})

	// Generate the code
	g.Execute()
}
