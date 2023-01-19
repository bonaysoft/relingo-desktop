package main

import (
	"context"
	"log"

	"github.com/bonaysoft/relingo-desktop/pkg/dal/model"
	"github.com/bonaysoft/relingo-desktop/pkg/dal/query"
	"github.com/bonaysoft/relingo-desktop/pkg/proxy"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx context.Context

	query *query.Query
	p     *proxy.Proxy
}

// NewApp creates a new App application struct
func NewApp() *App {
	gormdb, err := gorm.Open(sqlite.Open("gorm.db"))
	if err != nil {
		log.Fatalln(err)
	}

	gormdb.AutoMigrate(model.Word{})

	q := query.Use(gormdb)
	return &App{
		query: q,
		p:     proxy.NewProxy(q),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go a.p.Run()
}

func (a *App) shutdown(ctx context.Context) {
	a.p.Stop(ctx)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) []*model.Word {
	words, err := a.query.Word.Find()
	if err != nil {
		return nil
	}

	return words
}
