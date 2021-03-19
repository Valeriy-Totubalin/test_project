package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLink(t *testing.T) {
	id := 42
	login := "test_login"

	link := &Link{
		Id:        id,
		UserLogin: login,
	}

	assert.Equal(t, id, link.Id)
	assert.Equal(t, login, link.UserLogin)
}
