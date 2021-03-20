package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	linkTTL  time.Duration
	tokenTTL time.Duration
}

func NewConfig() *Config {
	return &Config{
		linkTTL:  getEnvAsHours("LINK_TTL", 24*time.Hour),
		tokenTTL: getEnvAsMinutes("TOKEN_TTL", 15*time.Minute),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsHours(name string, defaultVal time.Duration) time.Duration {
	valStr := getEnv(name, "")
	if value, err := strconv.Atoi(valStr); err == nil {
		return time.Duration(value) * time.Hour
	}

	return defaultVal
}

func getEnvAsMinutes(name string, defaultVal time.Duration) time.Duration {
	valStr := getEnv(name, "")
	if value, err := strconv.Atoi(valStr); err == nil {
		return time.Duration(value) * time.Minute
	}

	return defaultVal
}

func (c *Config) GetLinkTTL() *time.Duration {
	return &c.linkTTL
}

func (c *Config) GetTokenTTL() *time.Duration {
	return &c.tokenTTL
}
