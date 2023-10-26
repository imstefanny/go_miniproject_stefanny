package dto

type CreateTicketRequest struct {
	SeatID				uint			`json:"seat_id"`
	ShowID				uint			`json:"show_id"`
}
