package notifications

type Notification struct {
	ID          string `json:"id"`
	Context     string `json:"context"`
	MessageID   string `json:"message_id"`
	RecipientID string `json:"recipient_id"`
	SentAt      string `json:"sent_at"`
}
