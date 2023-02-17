package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/bonaysoft/relingo-desktop/pkg/dal/model"
	"github.com/bonaysoft/relingo-desktop/pkg/dal/query"
	"github.com/bonaysoft/relingo-desktop/pkg/proxy"
	"github.com/bonaysoft/relingo-desktop/pkg/relingo"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gen"

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

	q := query.Use(gormdb.Debug())
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

func (a *App) DownloadCert() error {
	opts := runtime.SaveDialogOptions{
		DefaultFilename: "ca.pem",
	}
	v, err := runtime.SaveFileDialog(a.ctx, opts)
	if err != nil {
		return err
	}

	src, err := os.Open("ca.pem")
	if err != nil {
		return err
	}
	dist, err := os.Create(v)
	if err != nil {
		return err
	}

	_, err = io.Copy(dist, src)
	return err
}

func (a *App) FindNewWords(q string) []*model.Word {
	date := func(t time.Time) time.Time {
		y, h, d := t.Date()
		return time.Date(y, h, d, 0, 0, 0, 0, time.Local)
	}
	today := time.Now()
	tomorrow := today.Add(time.Hour * 24)
	yesterday := today.Add(-time.Hour * 24)

	condsMap := map[string]gen.Condition{
		"today":     a.query.Word.UpdatedAt.Between(date(today), date(tomorrow)),
		"yesterday": a.query.Word.UpdatedAt.Between(date(yesterday), date(today)),
		"week":      a.query.Word.UpdatedAt.Between(date(GetSundayOfCurrentWeek(today)), date(tomorrow)),
	}
	conds := make([]gen.Condition, 0)
	cond, ok := condsMap[q]
	if ok {
		conds = append(conds, cond)
	}

	conds = append(conds, a.query.Word.Mastered.Is(false))
	words, err := a.query.Word.Where(conds...).Find()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// 查单词
	// 差词根
	// 查前缀

	return words
}

func (a *App) SubmitVocabulary(words []string) error {
	if err := a.rc.SubmitVocabulary(words); err != nil {
		return err
	}

	rows, err := a.query.Word.Where(a.query.Word.Source.In(words...)).Find()
	if err != nil {
		return err
	}

	for _, row := range rows {
		row.Mastered = true
	}

	return a.query.Word.Where(a.query.Word.Source.In(words...)).Save(rows...)
}

func (a *App) getBinds() []interface{} {
	return []interface{}{
		a,
		a.rc,
	}
}

func GetSundayOfCurrentWeek(t time.Time) time.Time {
	// 获取当前周的周一
	var offset int
	if t.Weekday() == time.Sunday {
		offset = 7
	} else {
		offset = int(t.Weekday())
	}
	return t.AddDate(0, 0, -offset)
}
