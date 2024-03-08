package collections

import (
	"fmt"
	"time"

	"github.com/charmingruby/upl/internal/core"
	"github.com/charmingruby/upl/internal/validation"
)

const (
	managerRole = "manager"
	defaultRole = "member"
)

func NewCollectionMember(role, accountID, collectionID string) (*CollectionMember, error) {
	member := CollectionMember{
		ID:             core.NewId(),
		Role:           defaultRole,
		UploadQuantity: 0,
		AccountID:      accountID,
		CollectionID:   collectionID,
		JoinedAt:       time.Now(),
		UpdatedAt:      nil,
		LeftAt:         nil,
	}

	if err := member.Validate(); err != nil {
		return nil, err
	}

	return &member, nil
}

type CollectionMember struct {
	ID             string     `json:"id"`
	Role           string     `db:"role" json:"role"`
	UploadQuantity int        `db:"upload_quantity" json:"upload_quantity"`
	AccountID      string     `json:"account_id"`
	CollectionID   string     `json:"collection_id"`
	JoinedAt       time.Time  `json:"joined_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	LeftAt         *time.Time `json:"left_at"`
}

func (m *CollectionMember) Validate() error {
	namedRole, err := m.validateRole(m.Role)
	if err != nil {
		return err
	}

	m.Role = namedRole

	if validation.IsEmpty(m.AccountID) {
		return &validation.ValidationError{
			Message: validation.NewRequiredFieldErrorMessage("account_id"),
		}
	}

	if validation.IsEmpty(m.CollectionID) {
		return &validation.ValidationError{
			Message: validation.NewRequiredFieldErrorMessage("collections_id"),
		}
	}

	return nil
}

func (m *CollectionMember) validateRole(role string) (string, error) {
	namedRole, ok := m.collectionMemberRoles()[role]

	if !ok {
		return "nil", fmt.Errorf("invalid role '%s'", role)
	}

	return namedRole, nil
}
func (a *CollectionMember) collectionMemberRoles() map[string]string {
	return map[string]string{
		managerRole: "manager",
		defaultRole: "member",
	}
}
