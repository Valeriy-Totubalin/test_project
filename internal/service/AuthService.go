package service

import (
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/config_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/pkg_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/repository_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/service_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/domain"
)

type AuthService struct {
	PasswordHasher pkg_interfaces.PasswordHasher
	UserRepository repository_interfaces.UserRepository
	TokenManager   pkg_interfaces.TokenManager
	Config         config_interfaces.GetterTokenTTL
}

func NewAuthService(
	userRepo repository_interfaces.UserRepository,
	passwordHasher pkg_interfaces.PasswordHasher,
	tokenManager pkg_interfaces.TokenManager,
	config config_interfaces.GetterTokenTTL,
) service_interfaces.AuthService {
	return &AuthService{
		UserRepository: userRepo,
		PasswordHasher: passwordHasher,
		TokenManager:   tokenManager,
		Config:         config,
	}
}

func (service *AuthService) SignUp(user *domain.User) error {
	passwordHash, err := service.PasswordHasher.GenerateHash(user.Password)
	if nil != err {
		return err
	}
	user.Password = passwordHash

	err = service.UserRepository.Create(user)
	if nil != err {
		return err
	}

	return nil
}

func (service *AuthService) SignIn(user *domain.User) (string, error) { // return token
	userReturned, err := service.UserRepository.GetByLogin(user.Login)
	if nil != err {
		return "", err
	}

	err = service.PasswordHasher.CheckPassword(user.Password, userReturned.Password)
	if nil != err {
		return "", err
	}

	token, err := service.TokenManager.NewJWT(userReturned.Id, service.Config.GetTokenTTL())
	if nil != err {
		return "", err
	}

	return token, nil
}

func (service *AuthService) IsExists(user *domain.User) bool {
	return service.UserRepository.IsExists(user)
}
