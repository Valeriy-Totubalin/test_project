package repository_interfaces

import "github.com/Valeriy-Totubalin/test_project/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	GetByLogin(login string) (*domain.User, error)
}
