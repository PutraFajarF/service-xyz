package entity

import "time"

type Loan struct {
	Id             int       `json:"id,omitempty"`
	ConsumerId     int       `json:"consumer_id,omitempty"`
	Amount         int       `json:"amount,omitempty"`
	TotalLoan      int       `json:"total_loan,omitempty"`
	LoanInterest   float64   `json:"load_interest,omitempty"`
	MonthlyPayment int       `json:"monthly_payment,omitempty"`
	Status         string    `json:"status,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}
