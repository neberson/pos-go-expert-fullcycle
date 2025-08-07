package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNeUser(t *testing.T) {
	user, err := NewUser("John Doe", "j0j.com", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "j0j.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "j0j.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.IsValidPassword("123456"))
	assert.False(t, user.IsValidPassword("1234567"))
	assert.NotEqual(t, "123456", user.Password)
}
