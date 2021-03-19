package link_manager

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/Valeriy-Totubalin/test_project/internal/domain"
	"github.com/dgrijalva/jwt-go"
)

type Manager struct {
	secret string
}

func NewManager(secret string) (*Manager, error) {
	if "" == secret {
		return nil, errors.New("empty secret key")
	}

	return &Manager{secret: secret}, nil
}

func (m *Manager) NewLink(link *domain.Link, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		Subject:   json.Marshal(link),
	})

	return token.SignedString([]byte(m.secret))
}
