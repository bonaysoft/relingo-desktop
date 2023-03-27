package main

import (
	"context"
	"log"
	"path/filepath"

	egClient "github.com/bonaysoft/engra/apis/client"
	"github.com/bonaysoft/relingo-desktop/pkg/config"
	"github.com/bonaysoft/relingo-desktop/pkg/dal"
	"github.com/bonaysoft/relingo-desktop/pkg/dal/model"
	"github.com/bonaysoft/relingo-desktop/pkg/dal/query"
	"github.com/bonaysoft/relingo-desktop/pkg/proxy"
	"github.com/bonaysoft/relingo-desktop/pkg/relingo"
	"github.com/bonaysoft/relingo-desktop/pkg/service"
	"github.com/bonaysoft/relingo-desktop/pkg/youdao"
)

// App struct
type App struct {
	ctx context.Context

	cm      *config.Manager
	query   *query.Query
	rProxy  *proxy.Proxy
	rc      *relingo.Client
	wordSvc *service.WordService
}

// NewApp creates a new App application struct
func NewApp() *App {
	cm := config.NewCM()
	cfg, err := cm.Load()
	if err != nil {
		log.Fatalln(err)
	}

	q, err := dal.Connect(filepath.Join(cm.GetConfigDir(), cfg.System.DBPath))
	if err != nil {
		log.Fatalln(err)
	}

	rc := relingo.NewClient(cfg.Relingo)
	egc := egClient.NewClient(cfg.System.EngraURL)
	return &App{
		cm:    cm,
		query: q,
		rc:    rc,

		rProxy:  proxy.NewProxy(q, rc),
		wordSvc: service.NewWordService(q, rc, egc),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go a.rProxy.Run()
}

func (a *App) shutdown(ctx context.Context) {
	_ = a.rProxy.Stop(ctx)
	_ = a.cm.Save()
}

func (a *App) getCtx() context.Context {
	return a.ctx
}

func (a *App) getBinds() []interface{} {
	return []interface{}{
		a.rc,
		a.wordSvc,
		service.NewSystem(a.getCtx),
		&model.Word{},
		youdao.NewClient(),
	}
}
