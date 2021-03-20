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
	UserRepository repository_interfaces.UserRepository
}

const ttlLink = 24 * time.Hour // время жизни ссылки 24 часа

func NewItemService(
	itemRepo repository_interfaces.ItemRepository,
	linkManager pkg_interfaces.LinkManager,
	userRepo repository_interfaces.UserRepository,
) service_interfaces.ItemService {
	return &ItemService{
		ItemRepository: itemRepo,
		LinkManager:    linkManager,
		UserRepository: userRepo,
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
	// добавить проверку: может ли этот пользователь передавать этот объект (является ли владельцем)
	libLink := &link_manager.Link{
		ItemId:    link.ItemId,
		UserLogin: link.UserLogin,
	}
	return service.LinkManager.NewLink(libLink, ttlLink)
}

func (service *ItemService) CanConfirm(tempLink string, userId int) (bool, error) {
	link, err := service.LinkManager.Parse(tempLink)
	if nil != err {
		return false, err
	}

	user, err := service.UserRepository.GetById(userId)
	if nil != err {
		return false, err
	}

	return user.Login == link.UserLogin, nil
}

func (service *ItemService) Confirm(tempLink string, userId int) error {
	link, err := service.LinkManager.Parse(tempLink)
	if nil != err {
		return err
	}

	// необходимо внутри вызвать в транзакции удаление и создание нового объекта
	err = service.ItemRepository.Transfer(link.ItemId, userId)
	if nil != err {
		return err
	}

	return nil
}
