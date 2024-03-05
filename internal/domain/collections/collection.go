package collections

func NewCollection() {}

type Collection struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Secret          string `json:"secret"`
	UploadsQuantity uint   `json:"uploads_quantity"`
	MembersQuantity uint   `json:"members_quantity"`
	TagID           string `json:"tag_id"`
	CreatorID       string `json:"creator_id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
