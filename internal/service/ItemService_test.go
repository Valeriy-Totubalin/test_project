package service

import (
	"testing"
	"time"

	"github.com/Valeriy-Totubalin/test_project/internal/domain"
	"github.com/Valeriy-Totubalin/test_project/pkg/link_manager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type LinkManagerMock struct {
	mock.Mock
}

func (m *LinkManagerMock) NewLink(link *link_manager.Link, ttl time.Duration) (string, error) {
	m.Called(link, ttl)

	return "temp_link", nil
}

func (m *LinkManagerMock) Parse(tempLink string) (*link_manager.Link, error) {
	m.Called(tempLink)

	return nil, nil
}

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
	linkManager := new(LinkManagerMock)

	serviceExpected := &ItemService{
		ItemRepository: repository,
		LinkManager:    linkManager,
	}

	serviceEqual := NewItemService(repository, linkManager)

	assert.Equal(t, serviceExpected, serviceEqual)
}

func TestCreate(t *testing.T) {
	repository := new(ItemRepositoryMock)
	linkManager := new(LinkManagerMock)

	service := NewItemService(repository, linkManager)

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
	linkManager := new(LinkManagerMock)

	service := NewItemService(repository, linkManager)

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
	linkManager := new(LinkManagerMock)

	service := NewItemService(repository, linkManager)

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

func TestGetTempLink(t *testing.T) {
	repository := new(ItemRepositoryMock)
	linkManager := new(LinkManagerMock)

	service := NewItemService(repository, linkManager)

	link := &domain.Link{
		ItemId:    42,
		UserLogin: "test_login",
	}
	libLink := &link_manager.Link{
		ItemId:    42,
		UserLogin: "test_login",
	}

	tempLink := "temp_link"

	linkManager.On("NewLink", libLink, 24*time.Hour).Return(tempLink, nil).Once()

	linkReturned, err := service.GetTempLink(link)

	linkManager.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, tempLink, linkReturned)
}
