package config_interfaces

import "time"

type GetterLinkTTL interface {
	GetLinkTTL() *time.Duration
}
