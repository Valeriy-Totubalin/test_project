package domain

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	id := 42
	login := "test_login"
	password := "password_string"

	user := User{
		Id: id,
		Login: login,
		Password: password,
	}

	assert.Equal(t, id, user.Id)
	assert.Equal(t, login, user.Login)
	assert.Equal(t, password, user.Password)
}