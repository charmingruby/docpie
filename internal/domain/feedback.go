package domain

import "time"

type Feedback struct {
	ID         string    `json:"id"`
	Rate       uint      `json:"rate"`
	Comment    string    `json:"comment"`
	CategoryID string    `json:"category_id"`
	ProductID  string    `json:"product_id"`
	AccountID  string    `json:"account_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type FeedbackUseCase interface {
	FetchFeedbacks(categoryID, productID, accountID string) ([]Feedback, error)
	GetFeedback(feedbackID string) (Feedback, error)
	CreateFeedback(rate, comment, categoryID, productID, accountID string) error
	DeleteFeedback(feedbackID, accountID string) error
}
