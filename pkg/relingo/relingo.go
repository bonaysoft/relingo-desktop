package relingo

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/samber/lo"
)

type RD interface {
	RespUserInfo | RespParseContent2 | Vocabulary | []VocabularyListItem | []DictItem
}

type Response[Data RD] struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    Data        `json:"data"`
}

func NewResponse[Data RD](data Data) *Response[Data] {
	return &Response[Data]{Data: data}
}

type RespUserInfo struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

type VocabularyBody struct {
	Id    string   `json:"id"`
	Type  string   `json:"type"`
	Words []string `json:"words,omitempty"`
}

type Vocabulary struct {
	Id        string    `json:"_id"`
	Uid       string    `json:"uid"`
	Id1       string    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Words     []string  `json:"words"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt int64     `json:"updatedAt"`
}

type VocabularyListItem struct {
	Count     int       `json:"count,omitempty"`
	Mastered  int       `json:"mastered,omitempty"`
	Id        string    `json:"_id,omitempty"`
	Uid       string    `json:"uid,omitempty"`
	Name      string    `json:"name"`
	Id1       string    `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt int64     `json:"updatedAt,omitempty"`
	Privilege string    `json:"privilege,omitempty"`
	Scope     string    `json:"scope,omitempty"`
	Level     string    `json:"level,omitempty"`
	Words     []string  `json:"words,omitempty"`
}

type RespParseContent2 struct {
	Words []Word `json:"words"`
}

type Word struct {
	Id            string   `json:"_id"`
	Phonetic      []string `json:"phonetic"`
	Variant       []string `json:"variant"`
	WordFrequency int      `json:"wordFrequency"`
	Source        string   `json:"source"`
	Display       *string  `json:"display,omitempty"`
	Translations  []struct {
		Target string  `json:"target"`
		Pos    string  `json:"pos"`
		Score  float64 `json:"score"`
	} `json:"translations"`
	Lang      string `json:"lang"`
	Mastered  bool   `json:"mastered"`
	Stared    bool   `json:"stared"`
	Sentences []struct {
		Id        string `json:"_id"`
		Uid       string `json:"uid"`
		Word      string `json:"word"`
		Sentence  string `json:"sentence"`
		Url       string `json:"url"`
		CreatedAt int64  `json:"createdAt"`
		UpdatedAt int64  `json:"updatedAt"`
		Md5       string `json:"md5"`
	} `json:"sentences"`
	Revision   bool   `json:"revision"`
	NeedRevise bool   `json:"needRevise"`
	Scope      string `json:"scope,omitempty"`
}

func (w *Word) String() string {
	b, _ := json.Marshal(w)
	return string(b)
}

type LockupBody struct {
	Text string `json:"text"`
	To   string `json:"to"`
}

type DictItem struct {
	Phonetic      []string      `json:"phonetic"`
	Variant       []string      `json:"variant"`
	WordFrequency int           `json:"wordFrequency"`
	Definition    string        `json:"definition"`
	Id            string        `json:"_id"`
	Source        string        `json:"source"`
	Lang          string        `json:"lang"`
	Translations  []Translation `json:"translations"`
	Display       string        `json:"display"`
	Mastered      bool          `json:"mastered"`
	Stared        bool          `json:"stared"`
	// Sentences  []interface{} `json:"sentences"`
	Revision   bool `json:"revision"`
	NeedRevise bool `json:"needRevise"`
}

func (d *DictItem) GetMeanings() string {
	meanings := lo.Map(d.Translations, func(item Translation, index int) string { return item.Target })
	return strings.Join(meanings, "；")
}

type Translation struct {
	Target string  `json:"target"`
	Pos    string  `json:"pos"`
	Score  float64 `json:"score"`
}
