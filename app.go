package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	egClient "github.com/bonaysoft/engra/apis/client"
	egModel "github.com/bonaysoft/engra/apis/graph/model"
	"github.com/bonaysoft/relingo-desktop/pkg/config"
	"github.com/bonaysoft/relingo-desktop/pkg/dal/model"
	"github.com/bonaysoft/relingo-desktop/pkg/dal/query"
	"github.com/bonaysoft/relingo-desktop/pkg/proxy"
	"github.com/bonaysoft/relingo-desktop/pkg/relingo"
	"github.com/bonaysoft/relingo-desktop/pkg/youdao"
	"github.com/samber/lo"
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
	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
		return nil
	}

	gormdb, err := gorm.Open(sqlite.Open(cfg.GetDBPath()))
	if err != nil {
		log.Fatalln(err)
	}

	gormdb.AutoMigrate(model.Word{})

	q := query.Use(gormdb.Debug())
	rc := relingo.NewClient()
	rc.TokenHook(cfg.SaveToken)
	rc.SetToken(cfg.GetToken())
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
		DefaultFilename: "relingo-desktop.crt",
	}
	v, err := runtime.SaveFileDialog(a.ctx, opts)
	if err != nil {
		return err
	}

	dist, err := os.Create(v)
	if err != nil {
		return err
	}

	_, err = io.Copy(dist, bytes.NewReader(proxy.GetCert()))
	return err
}

type Word model.Word

type ListResult struct {
	Items []Word `json:"items"`
	Total int64  `json:"total"`
}

func (a *App) FindNewWords(q string, pageNo, pageSize int) *ListResult {
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
	qd := a.query.Word.Where(conds...)
	total, err := qd.Count()
	words, err := qd.Offset((pageNo - 1) * pageSize).Limit(pageSize).Order(a.query.Word.Exposures.Desc()).Find()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	egc := egClient.NewClient("http://localhost:8081/query")
	resp, err := egClient.Lookup(egc.WithContext(context.Background()), lo.Map(words, func(item *model.Word, index int) string { return item.Name }))
	egvMaps := lo.SliceToMap(resp.Vocabularies, func(item egModel.Vocabulary) (string, egModel.Vocabulary) {
		return item.Name, item
	})

	items := make([]Word, 0)
	for _, word := range words {
		if v, ok := egvMaps[word.Name]; ok {
			word.EngraData = v
		}

		word.RawObject = word.ParseRawJSON()
		items = append(items, Word(*word))
	}

	return &ListResult{Items: items, Total: total}
}

func (a *App) SubmitVocabulary(words []string) error {
	if err := a.rc.SubmitVocabulary(words); err != nil {
		return err
	}

	rows, err := a.query.Word.Where(a.query.Word.Name.In(words...)).Find()
	if err != nil {
		return err
	}

	for _, row := range rows {
		row.Mastered = true
	}

	return a.query.Word.Where(a.query.Word.Name.In(words...)).Save(rows...)
}

func (a *App) getBinds() []interface{} {
	return []interface{}{
		a,
		a.rc,
		&model.Word{},
		youdao.NewClient(),
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
