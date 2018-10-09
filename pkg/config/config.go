package config

import (
	"fmt"
	"os"
)

func envDefault(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultValue
	}
	return val
}

// DBHost returns the POSTGRES_HOST if exists, otherwise localhost
func DBHost() string {
	return envDefault("POSTGRES_HOST", "localhost")
}

// DBName returns the POSTGRES_DB if exists, otherwise postgres
func DBName() string {
	return envDefault("POSTGRES_DB", "postgres")
}

// DBUser returns the POSTGRES_USER if exists, otherwise postgres
func DBUser() string {
	return envDefault("POSTGRES_USER", "postgres")
}

// DBPassword returns the POSTGRES_PASSWORD if exists, otherwise ""
func DBPassword() string {
	return envDefault("POSTGRES_PASSWORD", "")
}

// DBConnectionString forms & returns the DBConnectionString
func DBConnectionString() string {
	return fmt.Sprintf("host=%s dbname=%s user=%s password=%s",
		DBHost(),
		DBName(),
		DBUser(),
		DBPassword(),
	)
}
