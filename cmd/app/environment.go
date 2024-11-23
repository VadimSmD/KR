package main

import (
	"fmt"
	"os"

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
	pg := Postgres{DSN: os.Getenv("PG_URL")}
	env := environment{Postgres: pg}
	if err := cleanenv.ReadEnv(&env); err != nil {
		return environment{}, fmt.Errorf("newEnvironment: %w", err)
	}
	return env, nil
}
