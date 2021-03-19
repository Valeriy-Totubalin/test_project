package pkg_interfaces

import (
	"time"

	"github.com/Valeriy-Totubalin/test_project/internal/domain"
)

type LinkManager interface {
	NewLink(link *domain.Link, ttl time.Duration) (string, error)
	Parse(accessToken string) (*domain.Link, error) // возвращает id пользователя
}
