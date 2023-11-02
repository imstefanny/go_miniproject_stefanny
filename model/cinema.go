package model

import "time"

type Cinema struct {
	ID			 	uint			`json:"id"`
	Name			string		`json:"name"`
	Location  string		`json:"location"`
	Street   	string 		`json:"street"`
	City	 		string 		`json:"city"`
	Contact		string 		`json:"contact_info"`
	Studio		[]Studio
	CreatedAt	time.Time `json:"created_at"`
	UpdatedAt	time.Time `json:"updated_at"`
}
