// Package config to db
package config

import (
	"fmt"

	"github.com/caarlos0/env/v7"
)

// Config struct used to declare db connection
type Config struct {
	USER        string `env:"USER" envDefault:"postgres"`
	PostgresURL string `env:"pUrl" envDefault:"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"`
	PASSWORD    string `env:"PASSWORD" envDefault:"postgres"`
	PORT        int    `env:"PORT" envDefault:"5432"`
	DB          string `env:"DB" envDefault:"postgres"`
	MongoURL    string `env:"mUrl" envDefault:"mongodb://mongo:mongo@localhost:27017"`
	MongoDB     string `env:"DB" envDefault:"mongo"`
}

// NewConfig used to init config to db
func NewConfig() (*Config, error) {
	Cfg := &Config{}
	if err := env.Parse(Cfg); err != nil {
		return nil, fmt.Errorf("config - NewConfig: %v", err)
	}

	return Cfg, nil
}
