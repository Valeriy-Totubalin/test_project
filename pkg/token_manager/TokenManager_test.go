package token_manager

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewManager(t *testing.T) {
	secret := "secret_key"

	managerExpected := &Manager{
		secret: secret,
	}

	managerEqual, err := NewManager(secret)

	assert.Nil(t, err)
	assert.Equal(t, managerExpected, managerEqual)
}

func TestNewJWT(t *testing.T) {
	secret := "secret_key"
	manager, _ := NewManager(secret)

	token, err := manager.NewJWT(1, 15*time.Minute)

	assert.Nil(t, err)
	assert.Equal(t, reflect.String, reflect.TypeOf(token).Kind())
}

func TestParse(t *testing.T) {
	secret := "secret_key"
	userId := 1
	manager, _ := NewManager(secret)

	token, _ := manager.NewJWT(userId, 15*time.Minute)

	userIdResult, err := manager.Parse(token)

	assert.Nil(t, err)
	assert.Equal(t, userId, userIdResult)
}
