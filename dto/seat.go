package dto

type CreateSeatRequest struct {
	SeatNo				string		`json:"seat_no"`
	StudioID			uint			`json:"studio_id"`
}