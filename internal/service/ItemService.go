package service

import (
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/repository_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/service_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/domain"
)

type ItemService struct {
	ItemRepository repository_interfaces.ItemRepository
}

func NewItemService(
	itemRepo repository_interfaces.ItemRepository,
) service_interfaces.ItemService {
	return &ItemService{
		ItemRepository: itemRepo,
	}
}

func (service *ItemService) Create(item *domain.Item) error {
	return service.ItemRepository.Create(item)
}

func (service *ItemService) Delete(item *domain.Item) error {
	return service.ItemRepository.DeleteById(item.Id)
}

func (service *ItemService) GetAll() ([]*domain.Item, error) {
	items, err := service.ItemRepository.GetAll()
	if nil != err {
		return nil, err
	}

	return items, nil
}
