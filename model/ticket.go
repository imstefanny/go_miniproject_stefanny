package model

import "time"

type Ticket struct {
	ID						uint			`json:"id"`
	SeatID				uint			`json:"seat_id"`
	ShowID				uint			`json:"show_id"`
	TransactionID	uint			`json:"transaction_id"`
	CreatedAt			time.Time `json:"created_at"`
	UpdatedAt			time.Time `json:"updated_at"`
}
