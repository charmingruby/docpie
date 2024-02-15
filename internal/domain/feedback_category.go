package domain

import "time"

type FeedbackCategory struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ProductID   string    `json:"product_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type FeedbackCategoryUseCase interface {
	FetchFeedbackCategories()
	GetFeedbackCategory()
	CreateFeedbackCategory()
	UpdateFeedbackCategory()
	DeleteFeedbackCategory()
}
