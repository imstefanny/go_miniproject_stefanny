package model

import "time"

type Transaction struct {
	ID						uint			`json:"id"`
	AccountID			uint			`json:"account_id"`
	Tickets				[]Ticket	`gorm:"foreignKey:TransactionID"`
	Date					time.Time	`json:"date"`
	TotalPrice		int				`json:"total_price"`
	Status				string		`json:"status"`
	TicketCode		string		`json:"ticket_code"`
	CreatedAt			time.Time `json:"created_at"`
	UpdatedAt			time.Time `json:"updated_at"`
}