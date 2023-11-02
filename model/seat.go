package model

import "time"

type Seat struct {
	ID						uint			`json:"id"`
	SeatNo				string		`json:"seat_no"`
	StudioID			uint			`json:"studio_id"`
	CreatedAt			time.Time `json:"created_at"`
	UpdatedAt			time.Time `json:"updated_at"`
}
