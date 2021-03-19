package service_interfaces

import (
	"github.com/Valeriy-Totubalin/test_project/internal/domain"
)

type AuthService interface {
	SignUp(user *domain.User) error
	SignIn(user *domain.User) (string, error) // return token
	LogOut(userId int) error
}
