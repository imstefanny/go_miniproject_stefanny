package model

import "time"

type Studio struct {
	ID					uint			`json:"id"`
	Name				string		`json:"name"`
	Capacity		int				`json:"capacity"`
	CinemaID		uint			`json:"cinema_id"`
	CreatedAt		time.Time `json:"created_at"`
	UpdatedAt		time.Time `json:"updated_at"`
}