package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateItem(t *testing.T) {
	id := 42
	name := "item_name"

	item := Item{
		Id:   id,
		Name: name,
	}

	assert.Equal(t, id, item.Id)
	assert.Equal(t, name, item.Name)
}
