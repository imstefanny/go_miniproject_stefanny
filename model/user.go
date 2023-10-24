package model

import "time"

type User struct {
	ID			 		uint			`json:"id"`
	FirstName		string		`json:"first_name"`
	LastName	  string		`json:"last_name"`
	Email		   	string 		`json:"email"`
	Phone	 			string 		`json:"phone"`
	Address			string 		`json:"address"`
	Gender			string		`json:"gender"`
	DateOfBirth	time.Time	`json:"date_of_birth"`
	Pin					string		`json:"pin"`
	CreatedAt		time.Time `json:"created_at"`
	UpdatedAt		time.Time `json:"updated_at"`
}
