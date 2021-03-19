package service

import (
	"testing"

	"github.com/Valeriy-Totubalin/test_project/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ItemRepositoryMock struct {
	mock.Mock
}

func (m *ItemRepositoryMock) Create(item *domain.Item) error {
	m.Called(item)

	return nil
}

func (m *ItemRepositoryMock) DeleteById(itemId int) error {
	m.Called(itemId)

	return nil
}

func (m *ItemRepositoryMock) GetAll() ([]*domain.Item, error) {
	m.Called()

	return []*domain.Item{
		{
			Id: 42,
		},
		{
			Id: 23,
		},
		{
			Id: 97,
		},
	}, nil
}

func TestNewItemService(t *testing.T) {
	repository := new(ItemRepositoryMock)

	serviceExpected := &ItemService{
		ItemRepository: repository,
	}

	serviceEqual := NewItemService(repository)

	assert.Equal(t, serviceExpected, serviceEqual)
}

func TestCreate(t *testing.T) {
	repository := new(ItemRepositoryMock)
	service := NewItemService(repository)

	item := &domain.Item{
		Name: "item_name",
	}

	repository.On("Create", item).Return(nil).Once()

	err := service.Create(item)

	repository.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	repository := new(ItemRepositoryMock)
	service := NewItemService(repository)

	item := &domain.Item{
		Id: 42,
	}

	repository.On("DeleteById", item.Id).Return(nil).Once()

	err := service.Delete(item)

	repository.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestGetAll(t *testing.T) {
	repository := new(ItemRepositoryMock)
	service := NewItemService(repository)

	items := []*domain.Item{
		{
			Id: 42,
		},
		{
			Id: 23,
		},
		{
			Id: 97,
		},
	}

	repository.On("GetAll").Return(items).Once()

	itemsReturned, err := service.GetAll()

	repository.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, items, itemsReturned)
}
