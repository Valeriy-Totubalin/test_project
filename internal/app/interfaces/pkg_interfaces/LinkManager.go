package pkg_interfaces

import (
	"time"

	"github.com/Valeriy-Totubalin/test_project/pkg/link_manager"
)

type LinkManager interface {
	NewLink(link *link_manager.Link, ttl time.Duration) (string, error)
	Parse(tempLink string) (*link_manager.Link, error) // возвращает объект ссылки
}
