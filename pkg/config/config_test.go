package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func withEnvSet(key, value string, f func(key string)) {
	os.Setenv(key, value)
	f(key)
	os.Unsetenv(key)
}

func TestEnvBool(t *testing.T) {
	withEnvSet("BOOLEAN_KEY", "wrong", func(key string) {
		val := envBool(key, true)
		assert.False(t, val)

		val = envBool(key, false)
		assert.False(t, val)
	})

	withEnvSet("BOOLEAN_KEY", "TRUE", func(key string) {
		val := envBool(key, false)
		assert.True(t, val)

		val = envBool(key, true)
		assert.True(t, val)
	})

	// No Value set
	val := envBool("BOOLEAN_KEY", false)
	assert.False(t, val)

	val = envBool("BOOLEAN_KEY", true)
	assert.True(t, val)
}

func TestEnvInt(t *testing.T) {
	withEnvSet("MYINT", "10000", func(key string) {
		val, err := envInt(key, 3000)
		assert.NoError(t, err)
		assert.Equal(t, 10000, val)
	})

	withEnvSet("MYINT", "", func(key string) {
		val, err := envInt(key, 1000)
		assert.NoError(t, err)
		assert.Equal(t, 1000, val)
	})
}
