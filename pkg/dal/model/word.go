package model

import (
	"encoding/json"
	"time"

	"github.com/bonaysoft/engra/apis/graph/model"
	"github.com/bonaysoft/relingo-desktop/pkg/relingo"
)

type Word struct {
	Id        int              `json:"id"`
	Name      string           `json:"name" gorm:"column:source"`
	Exposures int              `json:"exposures"`
	Mastered  bool             `json:"mastered"`
	RawJSON   string           `json:"raw_json"`
	RawObject relingo.DictItem `json:"raw_object" gorm:"-"`
	EngraData model.Vocabulary `json:"engra_data" gorm:"-"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt time.Time        `json:"deleted_at"`
}

func (w *Word) ParseRawJSON() relingo.DictItem {
	var di relingo.DictItem
	json.Unmarshal([]byte(w.RawJSON), &di)
	return di
}
