package service_interfaces

import "github.com/Valeriy-Totubalin/test_project/internal/domain"

type ItemService interface {
	Create(item *domain.Item) error
	Delete(item *domain.Item) error
	GetAll() ([]*domain.Item, error)
}
