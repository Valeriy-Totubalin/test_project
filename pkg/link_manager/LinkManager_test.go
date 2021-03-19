package link_manager

import (
	"reflect"
	"testing"
	"time"

	"github.com/Valeriy-Totubalin/test_project/internal/domain"
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

func TestNewLink(t *testing.T) {
	secret := "secret_key"
	manager, _ := NewManager(secret)

	link = &domain.Link{
		Id:        42,
		UserLogin: "test_login",
	}

	tempLink, err := manager.NewLink(1, 15*time.Minute)

	assert.Nil(t, err)
	assert.Equal(t, reflect.String, reflect.TypeOf(tempLink).Kind())
}
