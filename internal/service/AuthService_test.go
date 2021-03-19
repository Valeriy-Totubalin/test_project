package service

import (
	"testing"
	"time"

	"github.com/Valeriy-Totubalin/test_project/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TokenManagerMock struct {
	mock.Mock
}

func (m *TokenManagerMock) NewJWT(userId int, ttl time.Duration) (string, error) {
	m.Called(userId, ttl)

	return "token", nil
}
func (m *TokenManagerMock) Parse(accessToken string) (int, error) { return 0, nil }

type PasswordHasherMock struct {
	mock.Mock
}

func (m *PasswordHasherMock) GenerateHash(password string) (string, error) {
	m.Called(password)

	return "password_hash", nil
}

func (m *PasswordHasherMock) CheckPassword(password string, hashedPassword string) error {
	m.Called(password, hashedPassword)

	return nil
}

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(user *domain.User) error {
	m.Called(user)

	return nil
}

func (m *UserRepositoryMock) GetByLogin(login string) (*domain.User, error) {
	m.Called(login)

	return &domain.User{
		Id:       1,
		Login:    "test_login",
		Password: "password_hash",
	}, nil
}

func TestNewAuthService(t *testing.T) {
	repository := new(UserRepositoryMock)
	passwordHasher := new(PasswordHasherMock)
	tokenManager := new(TokenManagerMock)

	serviceExpected := &AuthService{
		PasswordHasher: passwordHasher,
		UserRepository: repository,
		TokenManager:   tokenManager,
	}

	serviceEqual := NewAuthService(repository, passwordHasher, tokenManager)

	assert.Equal(t, serviceExpected, serviceEqual)
}

func TestSignUp(t *testing.T) {
	passwordHasher := new(PasswordHasherMock)
	repository := new(UserRepositoryMock)
	tokenManager := new(TokenManagerMock)
	service := NewAuthService(repository, passwordHasher, tokenManager)

	user := &domain.User{
		Login:    "test_login",
		Password: "test_password",
	}
	userWithHashedPassword := &domain.User{
		Login:    "test_login",
		Password: "password_hash",
	}

	passwordHasher.On("GenerateHash", user.Password).Once()
	repository.On("Create", userWithHashedPassword).Return(nil).Once()

	err := service.SignUp(user)

	passwordHasher.AssertExpectations(t)
	repository.AssertExpectations(t)

	assert.Equal(t, nil, err)
}

func TestSignIn(t *testing.T) {
	passwordHasher := new(PasswordHasherMock)
	repository := new(UserRepositoryMock)
	tokenManager := new(TokenManagerMock)
	service := NewAuthService(repository, passwordHasher, tokenManager)
	token := "token"

	user := &domain.User{
		Id:       1,
		Login:    "test_login",
		Password: "test_password",
	}

	userReturned := &domain.User{
		Id:       1,
		Login:    "test_login",
		Password: "password_hash",
	}

	repository.On("GetByLogin", user.Login).Return(userReturned).Once()
	passwordHasher.On("CheckPassword", user.Password, userReturned.Password).Once()
	tokenManager.On("NewJWT", user.Id, 15*time.Minute).Return(token).Once()

	tokenResult, err := service.SignIn(user)
	assert.Equal(t, token, tokenResult)

	passwordHasher.AssertExpectations(t)
	repository.AssertExpectations(t)
	tokenManager.AssertExpectations(t)

	assert.Equal(t, nil, err)
}
