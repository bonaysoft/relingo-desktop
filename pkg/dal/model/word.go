package model

import "time"

type Word struct {
	Id        int       `json:"id"`
	Source    string    `json:"source"`
	Exposures int       `json:"exposures"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
