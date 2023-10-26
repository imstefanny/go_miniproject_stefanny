package model

import "time"

type Show struct {
	ID						uint			`json:"id"`
	MovieID				uint			`json:"movie_id"`
	StudioID			uint			`json:"studio_id"`
	ShowDate			time.Time	`json:"show_date"`
	ShowStart			time.Time	`json:"show_start"`
	ShowEnd				time.Time	`json:"show_end"`
	Price					int				`json:"price"`
	CreatedAt			time.Time `json:"created_at"`
	UpdatedAt			time.Time `json:"updated_at"`
}