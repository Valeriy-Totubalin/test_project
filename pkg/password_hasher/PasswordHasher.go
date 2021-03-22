package password_hasher

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher struct {
}

func NewPasswordHasher() *PasswordHasher {
	return &PasswordHasher{}
}

func (service *PasswordHasher) GenerateHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (service *PasswordHasher) CheckPassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if nil != err {
		return errors.New("incorrect login or password")
	}

	return nil
}
