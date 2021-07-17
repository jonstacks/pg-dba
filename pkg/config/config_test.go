package config

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

func withEnvSet(key, value string, f func(key string)) {
	os.Setenv(key, value)
	f(key)
	os.Unsetenv(key)
}

func testAllBoolVariable(t *testing.T, key string, defaultValue bool, f func() bool) {
	for _, v := range []string{"true", "TRUE", "TrUE"} {
		withEnvSet(key, v, func(k string) {
			assert.True(t, f())
		})
	}

	for _, v := range []string{"false", "FALSE", "FaLSE"} {
		withEnvSet(key, v, func(k string) {
			assert.False(t, f())
		})
	}

	assert.Equal(t, defaultValue, f())
}

func testAllStringVariable(t *testing.T, key string, defaultValue string, f func() string) {
	assert.Equal(t, defaultValue, f())

	for _, v := range []string{"value one", "value_two", "VALUE_THREE"} {
		withEnvSet(key, v, func(k string) {
			assert.Equal(t, v, f())
		})
	}
}

func testAllTimeoutVariable(t *testing.T, key string, defaultValue int, f func() (int, error)) {
	val, err := f()
	assert.NoError(t, err)
	assert.Equal(t, defaultValue, val)

	for intVal, stringVal := range map[int]string{100: "100", 200: "200", 1000: "1000"} {
		withEnvSet(key, stringVal, func(k string) {
			val, err := f()
			assert.NoError(t, err)
			assert.Equal(t, intVal, val)
		})
	}

	withEnvSet(key, "SHOULD NOT PARSE", func(k string) {
		_, err := f()
		assert.Error(t, err)
	})
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

func TestEnvDefault(t *testing.T) {
	withEnvSet("TEST_VAR", "SomeValue", func(key string) {
		assert.Equal(t, "SomeValue", envDefault(key, "Default Value"))
	})

	assert.Equal(t, "Default Value", envDefault("TEST_VAR", "Default Value"))
}

func TestLogFormat(t *testing.T) {
	withEnvSet("LOG_FORMAT", "json", func(key string) {
		assert.IsType(t, &logrus.JSONFormatter{}, LogFormat())
	})

	withEnvSet("LOG_FORMAT", "text", func(key string) {
		assert.IsType(t, &logrus.TextFormatter{}, LogFormat())
	})

	assert.IsType(t, &logrus.TextFormatter{}, LogFormat())
}

func TestLogLevel(t *testing.T) {
	withEnvSet("LOG_LEVEL", "DEBUG", func(key string) {
		assert.Equal(t, logrus.DebugLevel, LogLevel())
	})

	withEnvSet("LOG_LEVEL", "WARN", func(key string) {
		assert.Equal(t, logrus.WarnLevel, LogLevel())
	})

	withEnvSet("LOG_LEVEL", "Error", func(key string) {
		assert.Equal(t, logrus.ErrorLevel, LogLevel())
	})

	withEnvSet("LOG_LEVEL", "INFO", func(key string) {
		assert.Equal(t, logrus.InfoLevel, LogLevel())
	})

	assert.Equal(t, logrus.InfoLevel, LogLevel())
}

func TestAnalyzeTimeoutSeconds(t *testing.T) {
	testAllTimeoutVariable(t, "ANALYZE_TIMEOUT_SECONDS", 600, AnalyzeTimeoutSeconds)
}

func TestFullVacuumTimeoutSeconds(t *testing.T) {
	testAllTimeoutVariable(t, "FULL_VACUUM_TIMEOUT_SECONDS", 600, FullVacuumTimeoutSeconds)
}

func TestVacuumTimeoutSeconds(t *testing.T) {
	testAllTimeoutVariable(t, "VACUUM_TIMEOUT_SECONDS", 600, VacuumTimeoutSeconds)
}

func TestBloatQueryTimeoutSeconds(t *testing.T) {
	testAllTimeoutVariable(t, "BLOAT_QUERY_TIMEOUT_SECONDS", 30, BloatQueryTimeoutSeconds)
}

func TestPostAnalyze(t *testing.T) {
	testAllBoolVariable(t, "POST_ANALYZE", true, PostAnalyze)
}

func TestPreAnalyze(t *testing.T) {
	testAllBoolVariable(t, "PRE_ANALYZE", true, PreAnalyze)
}

func TestVerbose(t *testing.T) {
	testAllBoolVariable(t, "VERBOSE", false, Verbose)
}

func TestDBHost(t *testing.T) {
	testAllStringVariable(t, "POSTGRES_HOST", "localhost", DBHost)
}

func TestDBUser(t *testing.T) {
	testAllStringVariable(t, "POSTGRES_USER", "postgres", DBUser)
}

func TestDBName(t *testing.T) {
	testAllStringVariable(t, "POSTGRES_DB", "postgres", DBName)
}

func TestDBPassword(t *testing.T) {
	testAllStringVariable(t, "POSTGRES_PASSWORD", `""`, DBPassword)
}

func TestDBConnectionString(t *testing.T) {
	withEnvSet("SSL_MODE", "require", func(key string) {
		assert.Equal(t, `host=localhost dbname=postgres user=postgres password="" sslmode=require`, DBConnectionString())
	})
}
