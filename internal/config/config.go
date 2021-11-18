package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DBHost         string `envconfig:"DB_HOST" default:"db"`
	DBPort         string `envconfig:"DB_PORT" default:"5432"`
	DBName         string `envconfig:"DB_NAME" default:"censys-local"`
	DBUser         string `envconfig:"DB_USER" default:"censys"`
	DBPassword     string `envconfig:"DB_PASSWORD" default:"censys"`
	DBSSLMode      string `envconfig:"DB_SSLMODE" default:"disable"` // disable, require, verify-ca or verify-full
}

// Load creates an instance of the config based on the environment
func Load(prefix string) (*Config, error) {
	c := &Config{}

	if err := envconfig.Process(prefix, c); err != nil {
		return nil, err
	}

	return c, nil
}