package factories_interfaces

import (
	"github.com/Valeriy-Totubalin/test_project/internal/repository/mysql/item_repository"
	"github.com/Valeriy-Totubalin/test_project/internal/repository/mysql/user_repository"
)

type RepositoriesFactory interface {
	NewUserRepository() *user_repository.UserRepository
	NewItemRepository() *item_repository.ItemRepository
}
