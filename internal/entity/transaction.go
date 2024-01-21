package entity

import "time"

type Transaction struct {
	Id          int       `json:"id,omitempty"`
	LoanId      int       `json:"loan_id,omitempty"`
	ConsumerId  int       `json:"consumer_id,omitempty"`
	Amount      int       `json:"amount,omitempty"`
	PaymentDate time.Time `json:"payment_date,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
