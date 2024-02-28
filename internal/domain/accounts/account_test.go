package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	validAccount, err := NewAccount(
		"john",
		"doe",
		"john@doe.com",
		"secret123",
	)

	assert.NoError(t, err)
	assert.NotEmpty(t, validAccount.ID)
	assert.NotEqual(t, "secret123", validAccount.Password)
	assert.Equal(t, defaultRole, validAccount.Role)

	validAccount.Role = managerRole
	assert.Equal(t, managerRole, validAccount.Role)
}

func TestAccountJSONSerialization(t *testing.T) {
	a, err := NewAccount(
		"john",
		"doe",
		"john@doe.com",
		"secret123",
	)

	assert.NoError(t, err)

	data, err := json.Marshal(a)
	assert.NoError(t, err)

	var newAccount Account
	err = json.Unmarshal(data, &newAccount)
	assert.NoError(t, err)

	assert.Equal(t, newAccount.ID, a.ID)
	assert.Equal(t, newAccount.Name, a.Name)
	assert.Equal(t, newAccount.LastName, a.LastName)
	assert.Equal(t, newAccount.Email, a.Email)
	assert.Equal(t, newAccount.AvatarURL, a.AvatarURL)
	assert.Equal(t, newAccount.Password, a.Password)
}
