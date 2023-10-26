package dto

import "time"

type CreateUserRequest struct {
	FirstName		string		`json:"first_name"`
	LastName	  string		`json:"last_name"`
	Email		   	string 		`json:"email"`
	Phone	 			string 		`json:"phone"`
	Address			string 		`json:"address"`
	Gender			string		`json:"gender"`
	DateOfBirth	time.Time	`json:"date_of_birth"`
	Pin					string		`json:"pin"`
	AccountID		uint			`json:"account_id"`
}