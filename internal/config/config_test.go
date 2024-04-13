package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	path := "../../.env"

	config, err := New(path)
	assert.NoError(t, err)

	assert.NotNil(t, config.HTTPConfig)
	assert.NotNil(t, config.PGConfig)
	assert.Equal(t, os.Getenv("JWT_SECRET_KEY"), config.JWT_SECRET)
	assert.Equal(t, os.Getenv("PASS_HASH_SECRET"), config.PASS_SECRET)
}
