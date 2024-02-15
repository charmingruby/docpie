package domain

import "time"

type Feedback struct {
	ID        string    `json:"id"`
	Rate      uint      `json:"rate"`
	Comment   string    `json:"comment"`
	ProductID string    `json:"product_id"`
	AccountID string    `json:"account_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
