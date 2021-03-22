package service

import (
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/config_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/pkg_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/repository_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/service_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/domain"
	"github.com/Valeriy-Totubalin/test_project/pkg/link_manager"
)

type ItemService struct {
	ItemRepository repository_interfaces.ItemRepository
	LinkManager    pkg_interfaces.LinkManager
	UserRepository repository_interfaces.GetterUserById
	Config         config_interfaces.GetterLinkTTL
}

func NewItemService(
	itemRepo repository_interfaces.ItemRepository,
	linkManager pkg_interfaces.LinkManager,
	userRepo repository_interfaces.GetterUserById,
	config config_interfaces.GetterLinkTTL,
) service_interfaces.ItemService {
	return &ItemService{
		ItemRepository: itemRepo,
		LinkManager:    linkManager,
		UserRepository: userRepo,
		Config:         config,
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

	return service.LinkManager.NewLink(libLink, service.Config.GetLinkTTL())
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

	err = service.ItemRepository.Transfer(link.ItemId, userId)
	if nil != err {
		return err
	}

	return nil
}
