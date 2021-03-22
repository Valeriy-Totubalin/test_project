package factories

import (
	"github.com/Valeriy-Totubalin/test_project/db/orm"
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/config_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/factories_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/repository/mysql/item_repository"
	"github.com/Valeriy-Totubalin/test_project/internal/repository/mysql/user_repository"
)

type RepositoriesFactory struct {
	config config_interfaces.DBConfig
}

func NewRepositoriesFactory(config config_interfaces.DBConfig) factories_interfaces.RepositoriesFactory {
	return &RepositoriesFactory{
		config: config,
	}
}

func (f *RepositoriesFactory) NewUserRepository() *user_repository.UserRepository {
	gorm := orm.NewGormDB(f.config)
	return user_repository.NewUserRepository(gorm)
}

func (f *RepositoriesFactory) NewItemRepository() *item_repository.ItemRepository {
	gorm := orm.NewGormDB(f.config)
	return item_repository.NewItemRepository(gorm)
}
