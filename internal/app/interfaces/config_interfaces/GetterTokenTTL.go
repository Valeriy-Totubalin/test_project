package config_interfaces

import "time"

type GetterTokenTTL interface {
	GetTokenTTL() time.Duration
}
