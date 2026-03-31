package config

import "os"

// Config holds application configuration.
type Config struct {
	Env  string
	Port string
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Env:  env,
		Port: port,
	}, nil
}
