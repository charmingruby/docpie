package domain

import "time"

type FeedbackCategory struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type FeedbackCategoryUseCase interface {
	CreateFeedbackCategory(name, description string) error
}
