package service

import (
	"context"
	"fmt"
	"time"

	egClient "github.com/bonaysoft/engra/apis/client"
	egModel "github.com/bonaysoft/engra/apis/graph/model"
	"github.com/bonaysoft/relingo-desktop/pkg/dal/model"
	"github.com/bonaysoft/relingo-desktop/pkg/dal/query"
	"github.com/bonaysoft/relingo-desktop/pkg/relingo"
	"github.com/samber/lo"
	"gorm.io/gen"
)

type Word model.Word

type ListResult struct {
	Items []Word `json:"items"`
	Total int64  `json:"total"`
}

type WordService struct {
	query *query.Query

	rc  *relingo.Client
	egc *egClient.Client
}

func NewWordService(q *query.Query, rc *relingo.Client, egc *egClient.Client) *WordService {
	return &WordService{query: q, rc: rc, egc: egc}
}

func (ws *WordService) FindNewWords(q string, pageNo, pageSize int) *ListResult {
	date := func(t time.Time) time.Time {
		y, h, d := t.Date()
		return time.Date(y, h, d, 0, 0, 0, 0, time.Local)
	}
	today := time.Now()
	tomorrow := today.Add(time.Hour * 24)
	yesterday := today.Add(-time.Hour * 24)

	condsMap := map[string]gen.Condition{
		"today":     ws.query.Word.UpdatedAt.Between(date(today), date(tomorrow)),
		"yesterday": ws.query.Word.UpdatedAt.Between(date(yesterday), date(today)),
		"week":      ws.query.Word.UpdatedAt.Between(date(getSundayOfCurrentWeek(today)), date(tomorrow)),
	}
	conds := make([]gen.Condition, 0)
	cond, ok := condsMap[q]
	if ok {
		conds = append(conds, cond)
	}

	conds = append(conds, ws.query.Word.Mastered.Is(false))
	qd := ws.query.Word.Where(conds...)
	total, err := qd.Count()
	words, err := qd.Offset((pageNo - 1) * pageSize).Limit(pageSize).Order(ws.query.Word.Exposures.Desc()).Find()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	resp, err := egClient.Lookup(ws.egc.WithContext(context.Background()), lo.Map(words, func(item *model.Word, index int) string { return item.Name }))
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

func (ws *WordService) GetRootByName(word string) (*egModel.Vocabulary, error) {
	resp, err := egClient.LookupRoot(ws.egc.WithContext(context.Background()), word)
	if err != nil {
		return nil, err
	}

	return &resp.VocabularyRootTree, nil
}

func (ws *WordService) SubmitVocabulary(words []string) error {
	if err := ws.rc.SubmitVocabulary(words); err != nil {
		return err
	}

	rows, err := ws.query.Word.Where(ws.query.Word.Name.In(words...)).Find()
	if err != nil {
		return err
	}

	for _, row := range rows {
		row.Mastered = true
	}

	return ws.query.Word.Where(ws.query.Word.Name.In(words...)).Save(rows...)
}

func getSundayOfCurrentWeek(t time.Time) time.Time {
	// 获取当前周的周一
	var offset int
	if t.Weekday() == time.Sunday {
		offset = 7
	} else {
		offset = int(t.Weekday())
	}
	return t.AddDate(0, 0, -offset)
}
