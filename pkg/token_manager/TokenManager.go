package token_manager

import (
	"errors"
	"strconv"
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

func (m *Manager) NewJWT(userId int, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		Subject:   strconv.Itoa(userId),
	})

	return token.SignedString([]byte(m.secret))
}

func (m *Manager) Parse(accessToken string) (int, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected string method")
		}
		return []byte(m.secret), nil
	})
	if nil != err {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("error get user claims from token")
	}

	userId, err := strconv.Atoi(claims["sub"].(string))
	if nil != err {
		return 0, err
	}

	return userId, nil
}
