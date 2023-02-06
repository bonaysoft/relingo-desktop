package model

import "time"

type Word struct {
	Id        int       `json:"id"`
	Source    string    `json:"source"`
	Exposures int       `json:"exposures"`
	Mastered  bool      `json:"mastered"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// Vocabulary 保存单词的词根词缀
type Vocabulary struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Root   string `json:"root"`
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

// RootsAffixes 词根词缀
type RootsAffixes struct {
	Name    string `json:"name"`
	Meaning string `json:"meaning"`
}
