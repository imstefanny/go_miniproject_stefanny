package dto

type CreateTransactionRequest struct {
	AccountID				uint			`json:"account_id"`
	ShowID					uint			`json:"show_id"`
	SeatID					[]uint		`json:"seat_id"`
}