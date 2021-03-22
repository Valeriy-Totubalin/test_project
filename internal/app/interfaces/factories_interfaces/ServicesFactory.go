package factories_interfaces

import "github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/service_interfaces"

type ServicesFactory interface {
	NewAuthService() (service_interfaces.AuthService, error)
	NewItemService() (service_interfaces.ItemService, error)
}
