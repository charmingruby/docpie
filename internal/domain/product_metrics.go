package domain

import "time"

type ProductMetrics struct {
	ID             string    `json:"id"`
	TotalFeedbacks uint      `json:"total_feedbacks"`
	AmountOfStars  uint      `json:"amount_of_stars"`
	ProductID      string    `json:"product_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ProductMetricsUseCase interface {
}
