package config_interfaces

type DBConfig interface {
	GetHost() string
	GetUser() string
	GetPassword() string
	GetPort() string
	GetName() string
}
