package model

import "time"

type Account struct {
	ID					uint			`json:"id"`
	Username		string		`json:"username"`
	Email				string		`json:"email"`
	Password		string		`json:"password"`
	Role				string		`json:"role"`
	UserID			User			`gorm:"foreignKey:AccountID"`
	CreatedAt		time.Time `json:"created_at"`
	UpdatedAt		time.Time `json:"updated_at"`
}

type AccountResponse struct {
	ID			 	uint			`json:"id"`
	Username		string		`json:"username"`
	Email    	string 		`json:"email"`
	Token 		string 		`json:"token"`
	CreatedAt	time.Time `json:"created_at"`
	UpdatedAt	time.Time `json:"updated_at"`
}