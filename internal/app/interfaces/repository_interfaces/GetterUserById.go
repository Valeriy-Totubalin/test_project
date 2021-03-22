package repository_interfaces

import "github.com/Valeriy-Totubalin/test_project/internal/domain"

type GetterUserById interface {
	GetById(userId int) (*domain.User, error)
}
