package config_interfaces

import "time"

type ServerConfig interface {
	GetPort() string
	GetReadTimeout() time.Duration
	GetWriteTimeout() time.Duration
}
