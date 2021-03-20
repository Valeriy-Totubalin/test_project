package service

import (
	"time"

	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/pkg_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/repository_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/service_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/domain"
	"github.com/Valeriy-Totubalin/test_project/pkg/link_manager"
)

type ItemService struct {
	ItemRepository repository_interfaces.ItemRepository
	LinkManager    pkg_interfaces.LinkManager
}

const ttlLink = 24 * time.Hour // время жизни ссылки 24 часа

func NewItemService(
	itemRepo repository_interfaces.ItemRepository,
	linkManager pkg_interfaces.LinkManager,
) service_interfaces.ItemService {
	return &ItemService{
		ItemRepository: itemRepo,
		LinkManager:    linkManager,
	}
}

func (service *ItemService) Create(item *domain.Item) error {
	return service.ItemRepository.Create(item)
}

func (service *ItemService) Delete(item *domain.Item) error {
	return service.ItemRepository.DeleteById(item.Id)
}

func (service *ItemService) GetAll() ([]*domain.Item, error) {
	return service.ItemRepository.GetAll()
}

func (service *ItemService) GetTempLink(link *domain.Link) (string, error) {
	libLink := &link_manager.Link{
		ItemId:    link.ItemId,
		UserLogin: link.UserLogin,
	}
	return service.LinkManager.NewLink(libLink, ttlLink)
}
