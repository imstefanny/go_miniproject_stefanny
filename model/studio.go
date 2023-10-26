package model

import "time"

type Studio struct {
	Name				string		`json:"name"`
	Capacity		int				`json:"capacity"`
	CinemaID		uint			`json:"cinema_id"`
	CreatedAt		time.Time `json:"created_at"`
	UpdatedAt		time.Time `json:"updated_at"`
}