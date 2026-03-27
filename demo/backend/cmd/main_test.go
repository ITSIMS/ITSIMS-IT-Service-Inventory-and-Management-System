package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv_ReturnsEnvValue(t *testing.T) {
	const key = "ITSIMS_TEST_KEY_12345"
	os.Setenv(key, "from_env")
	defer os.Unsetenv(key)

	assert.Equal(t, "from_env", getEnv(key, "default"))
}

func TestGetEnv_ReturnsDefault(t *testing.T) {
	const key = "ITSIMS_TEST_KEY_MISSING_99999"
	os.Unsetenv(key)

	assert.Equal(t, "fallback", getEnv(key, "fallback"))
}
