package main

import (
	"context"
	"log"

	"github.com/bonaysoft/relingo-desktop/pkg/dal/model"
	"github.com/bonaysoft/relingo-desktop/pkg/dal/query"
	"github.com/bonaysoft/relingo-desktop/pkg/proxy"
	"github.com/bonaysoft/relingo-desktop/pkg/relingo"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx context.Context

	query *query.Query
	p     *proxy.Proxy
	rc    *relingo.Client
}

// NewApp creates a new App application struct
func NewApp() *App {
	gormdb, err := gorm.Open(sqlite.Open("/Users/yanbo/CloudDocs/relingo-desktop/gorm.db"))
	if err != nil {
		log.Fatalln(err)
	}

	gormdb.AutoMigrate(model.Word{})

	q := query.Use(gormdb)
	rc := relingo.NewClient()
	return &App{
		query: q,
		rc:    rc,
		p:     proxy.NewProxy(q, rc),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go a.p.Run()
}

func (a *App) shutdown(ctx context.Context) {
	_ = a.p.Stop(ctx)
}

func (a *App) FindNewWords() []*model.Word {
	words, err := a.query.Word.Find()
	if err != nil {
		return nil
	}

	return words
}

func (a *App) getBinds() []interface{} {
	return []interface{}{
		a,
		a.rc,
	}
}
