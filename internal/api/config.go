package api

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/jsnfwlr/vexil/internal/db"
)

type Config struct {
	DBClient  *db.Client
	EnableSSE bool `env:"ENABLE_SSE"`
}

func LoadConfig(DBClient *db.Client) (cfg Config, fault error) {
	c := Config{}
	if err := env.Parse(&c); err != nil {
		return Config{}, fmt.Errorf("could not load config: %w", err)
	}

	c.DBClient = DBClient

	return c, nil
}
