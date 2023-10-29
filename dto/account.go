package dto

type CreateAccountRequest struct {
	Username		string		`json:"username"`
	Email				string		`json:"email"`
	Password		string		`json:"password"`
	Role				string		`json:"role"`
}

type LoginAccountRequest struct {
	Username		string		`json:"username"`
	Email				string		`json:"email"`
	Password		string		`json:"password"`
}
