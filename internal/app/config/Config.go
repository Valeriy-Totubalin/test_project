package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/config_interfaces"
	"github.com/joho/godotenv"
)

type Config struct {
	linkTTL     time.Duration
	tokenTTL    time.Duration
	db          config_interfaces.DBConfig
	server      config_interfaces.ServerConfig
	tokenSecret string
	linkSecret  string
}

func NewConfig(dbConfig config_interfaces.DBConfig, serverConfig config_interfaces.ServerConfig) *Config {
	return &Config{
		linkTTL:     getEnvAsHours("LINK_TTL", 24*time.Hour),
		tokenTTL:    getEnvAsMinutes("TOKEN_TTL", 15*time.Minute),
		db:          dbConfig,
		server:      serverConfig,
		tokenSecret: getEnv("TOKEN_SECRET", "zPOA4JnM1x1XOxGwo7kknGlcikudhMd09WnJ2DPWdi8"),
		linkSecret:  getEnv("LINK_SECRET", "agKfggjJJs1kfpdgk95jeihyA6gnsaJheiuJ2DP98wr"),
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

func getEnvAsSeconds(name string, defaultVal time.Duration) time.Duration {
	valStr := getEnv(name, "")
	if value, err := strconv.Atoi(valStr); err == nil {
		return time.Duration(value) * time.Second
	}

	return defaultVal
}

func (c *Config) GetLinkTTL() time.Duration {
	return c.linkTTL
}

func (c *Config) GetTokenTTL() time.Duration {
	return c.tokenTTL
}

func (c *Config) DB() config_interfaces.DBConfig {
	return c.db
}

func (c *Config) Srv() config_interfaces.ServerConfig {
	return c.server
}

func (c *Config) GetTokenSecret() string {
	return c.tokenSecret
}

func (c *Config) GetLinkSecret() string {
	return c.linkSecret
}

func Init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
