package service

import (
	"testing"

	"github.com/Valeriy-Totubalin/test_project/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type PasswordHasherMock struct {
	mock.Mock
}

func (m *PasswordHasherMock) GenerateHash(password string) (string, error) {
	m.Called(password)

	return "password_hash", nil
}

func (m *PasswordHasherMock) CheckPassword(password string, hashedPassword string) error { return nil }

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(user domain.User) error {
	m.Called(user)

	return nil
}

func (m *UserRepositoryMock) GetByLogin(login string) (*domain.User, error) { return nil, nil }

func TestNewAuthService(t *testing.T) {
	repository := new(UserRepositoryMock)
	passwordHasher := new(PasswordHasherMock)

	serviceExpected := &AuthService{
		PasswordHasher: passwordHasher,
		UserRepository: repository,
	}
	serviceEqual := NewAuthService(repository, passwordHasher)

	assert.Equal(t, serviceExpected, serviceEqual)
}

func TestSignUp(t *testing.T) {
	passwordHasher := new(PasswordHasherMock)
	repository := new(UserRepositoryMock)
	service := NewAuthService(repository, passwordHasher)

	user := domain.User{
		Login:    "test_login",
		Password: "test_password",
	}
	userWithHashedPassword := domain.User{
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
