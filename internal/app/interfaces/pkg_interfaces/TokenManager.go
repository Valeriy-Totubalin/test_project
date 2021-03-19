package pkg_interfaces

import "time"

type TokenManager interface {
	NewJWT(userId int, ttl time.Duration) (string, error)
	Parse(accessToken string) (int, error) // возвращает id пользователя
}
