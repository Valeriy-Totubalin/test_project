package factories

import (
	"github.com/Valeriy-Totubalin/test_project/internal/app/config"
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/factories_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/service_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/service"
	"github.com/Valeriy-Totubalin/test_project/pkg/link_manager"
	"github.com/Valeriy-Totubalin/test_project/pkg/password_hasher"
	"github.com/Valeriy-Totubalin/test_project/pkg/token_manager"
)

type ServicesFactory struct {
	config      config.Config
	repoFactory factories_interfaces.RepositoriesFactory
}

func NewServicesFactory(config config.Config) factories_interfaces.ServicesFactory {
	repoFactory := NewRepositoriesFactory(config.DB())

	return &ServicesFactory{
		config:      config,
		repoFactory: repoFactory,
	}
}

func (f *ServicesFactory) NewUserService() (service_interfaces.AuthService, error) {
	userRepo := f.repoFactory.NewUserRepository()
	passwordHasher := password_hasher.NewPasswordHasher()
	tokenManager, err := token_manager.NewManager(f.config.GetTokenSecret())

	if nil != err {
		return nil, err
	}

	userService := service.NewAuthService(
		userRepo,
		passwordHasher,
		tokenManager,
		f.config,
	)

	return userService, nil
}

func (f *ServicesFactory) NewItemService() (service_interfaces.ItemService, error) {
	itemRepo := f.repoFactory.NewItemRepository()
	linkManager, err := link_manager.NewManager(f.config.GetLinkSecret())
	userRepo := f.repoFactory.NewUserRepository()

	if nil != err {
		return nil, err
	}

	itemService := service.NewItemService(
		itemRepo,
		linkManager,
		userRepo,
		f.config,
	)

	return itemService, nil
}
