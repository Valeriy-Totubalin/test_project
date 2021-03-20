package link_manager

import (
	"encoding/json"
	"errors"
	"time"

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

// не уверен, что кастомные либы должны знать о домене. возврат json? ансериализация уже в сервисе
func (m *Manager) NewLink(link *Link, ttl time.Duration) (string, error) {
	jsonLink, err := json.Marshal(link)
	if nil != err {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		Subject:   string(jsonLink),
	})

	return token.SignedString([]byte(m.secret))
}

func (m *Manager) Parse(tempLink string) (*Link, error) {
	token, err := jwt.Parse(tempLink, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected string method")
		}
		return []byte(m.secret), nil
	})
	if nil != err {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("error get data claims from link")
	}

	jsonLink := claims["sub"].(string)
	link := &Link{}
	json.Unmarshal([]byte(jsonLink), link)

	return link, nil
}
