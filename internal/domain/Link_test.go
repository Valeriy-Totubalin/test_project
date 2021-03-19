package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLink(t *testing.T) {
	id := 42
	login := "test_login"

	link := &Link{
		ItemId:    id,
		UserLogin: login,
	}

	assert.Equal(t, id, link.ItemId)
	assert.Equal(t, login, link.UserLogin)
}
