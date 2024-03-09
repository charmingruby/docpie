package collections

import (
	"fmt"
	"time"

	"github.com/charmingruby/upl/internal/core"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/charmingruby/upl/internal/validation/errs"
)

const (
	managerRole = "manager"
	defaultRole = "member"
)

func NewCollectionMember(accountID, collectionID string) (*CollectionMember, error) {
	member := CollectionMember{
		ID:              core.NewId(),
		Role:            defaultRole,
		UploadsQuantity: 0,
		AccountID:       accountID,
		CollectionID:    collectionID,
		JoinedAt:        time.Now(),
		UpdatedAt:       nil,
		LeftAt:          nil,
	}

	if err := member.Validate(); err != nil {
		return nil, err
	}

	return &member, nil
}

type CollectionMember struct {
	ID              string     `db:"id" json:"id"`
	Role            string     `db:"role" json:"role"`
	UploadsQuantity int        `db:"uploads_quantity" json:"uploads_quantity"`
	AccountID       string     `db:"account_id" json:"account_id"`
	CollectionID    string     `db:"collection_id" json:"collection_id"`
	JoinedAt        time.Time  `db:"joined_at" json:"joined_at"`
	UpdatedAt       *time.Time `db:"updated_at" json:"updated_at"`
	LeftAt          *time.Time `db:"left_at" json:"left_at"`
}

func (m *CollectionMember) Validate() error {
	namedRole, err := m.validateRole(m.Role)
	if err != nil {
		return err
	}

	m.Role = namedRole

	if validation.IsEmpty(m.AccountID) {
		return &errs.ValidationError{
			Message: errs.EntitieisRequiredFieldErrorMessage("account_id"),
		}
	}

	if validation.IsEmpty(m.CollectionID) {
		return &errs.ValidationError{
			Message: errs.EntitieisRequiredFieldErrorMessage("collections_id"),
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
