package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateItem(t *testing.T) {
	id := 42
	name := "item_name"
	userId := 7

	item := Item{
		Id:     id,
		Name:   name,
		UserId: userId,
	}

	assert.Equal(t, id, item.Id)
	assert.Equal(t, name, item.Name)
	assert.Equal(t, userId, item.UserId)
}
