package dto

import "time"

type CreateShowRequest struct {
	MovieID				uint			`json:"movie_id"`
	StudioID			uint			`json:"studio_id"`
	ShowDate			time.Time	`json:"show_date"`
	ShowStart			time.Time	`json:"show_start"`
	ShowEnd				time.Time	`json:"show_end"`
	Price					int				`json:"price"`
}