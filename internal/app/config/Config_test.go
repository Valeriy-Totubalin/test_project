package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	configExpected := &Config{
		linkTTL:  24 * time.Hour,
		tokenTTL: 15 * time.Minute,
	}

	configEqual := NewConfig()

	assert.Equal(t, configExpected, configEqual)
}

func TestGetLinkTTL(t *testing.T) {
	config := NewConfig()
	ttl := config.GetLinkTTL()
	ttlWithExpectedType := time.Minute

	assert.IsType(t, &ttlWithExpectedType, ttl)
}

func TestGetTokenTTL(t *testing.T) {
	config := NewConfig()
	ttl := config.GetTokenTTL()
	ttlWithExpectedType := time.Hour

	assert.IsType(t, &ttlWithExpectedType, ttl)
}
