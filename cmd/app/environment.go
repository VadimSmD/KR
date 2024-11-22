package main

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	environment struct {
		Postgres Postgres
	}
	Postgres struct {
		DSN string `env:"PG_DSN" env-required:"true"`
	}
)

func newEnvironment() (environment, error) {
	var env environment
	if err := cleanenv.ReadEnv(&env); err != nil {
		return environment{}, fmt.Errorf("newEnvironment: %w", err)
	}
	return env, nil
}
