package dto

type CreateCinemaRequest struct {
	Name			string		`json:"name"`
	Location  string		`json:"location"`
	Street   	string 		`json:"street"`
	City	 		string 		`json:"city"`
	Contact		string 		`json:"contact_info"`
}