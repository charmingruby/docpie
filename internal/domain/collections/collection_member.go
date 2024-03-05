package collections

import "time"

type CollectionMember struct {
	ID           string     `json:"id"`
	AccountID    string     `json:"account_id"`
	CollectionID string     `json:"collection_id"`
	JoinedAt     time.Time  `json:"joined_at"`
	LeftAt       *time.Time `json:"left_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
