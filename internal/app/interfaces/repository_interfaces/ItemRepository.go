package repository_interfaces

import "github.com/Valeriy-Totubalin/test_project/internal/domain"

type ItemRepository interface {
	Create(item *domain.Item) error
	DeleteById(itemId int) error
	GetAll(userId int) ([]*domain.Item, error)
	Transfer(itemId int, userId int) error
	GetById(itemId int) (*domain.Item, error)
}
