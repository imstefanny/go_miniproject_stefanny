package model

import "time"

type Movie struct {
	ID			 			uint			`json:"id"`
	Title					string		`json:"title"`
	Duration  		int				`json:"duration"`
	ReleaseDate		time.Time	`json:"release_date"`
	Genre					string		`json:"genre"`
	Rating				float32		`json:"rating"`
	Synopsis			string		`json:"synopsis"`
	Producer			string		`json:"producer"`
	Director			string		`json:"director"`
	Writer				string		`json:"writer"`
	Cast					string		`json:"cast"`
	Distributor		string		`json:"distributor"`
	Type					string		`json:"type"`
	CreatedAt			time.Time `json:"created_at"`
	UpdatedAt			time.Time `json:"updated_at"`
}
