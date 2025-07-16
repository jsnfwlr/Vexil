package db

import (
	"fmt"
	"log/slog"

	env "github.com/caarlos0/env/v10"
)

type Secret string

func (Secret) LogValue() slog.Value {
	return slog.StringValue("REDACTED_SECRET")
}

type Config struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     string `env:"POSTGRES_PORT" envDefault:"5433"`
	Database string `env:"POSTGRES_DATABASE" envDefault:"vexil"` // Database name
	Username string `env:"POSTGRES_USER" envDefault:"vexil"`     // Username
	Password Secret `env:"POSTGRES_PASSWORD" envDefault:"vexil"` // Password
}

func (c Config) GetURI() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.Username, c.Password, c.Host, c.Port, c.Database)
}

func (c Config) GetRedactedURI() string {
	return fmt.Sprintf("postgres://%s:REDACTED@%s:%s/%s?sslmode=disable", c.Username, c.Host, c.Port, c.Database)
}

func (c Config) GetHost() string {
	return c.Host
}

func (c Config) GetPort() string {
	return c.Port
}

func (c Config) GetDatabase() string {
	return c.Database
}

func LoadConfig() (cfg Config, fault error) {
	c := Config{}
	if err := env.Parse(&c); err != nil {
		return Config{}, fmt.Errorf("could not load config: %w", err)
	}
	return c, nil
}
