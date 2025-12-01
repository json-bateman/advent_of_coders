package aoc

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port  int  `envconfig:"PORT" default:"4422"`
	IsDev bool `default:"false" split_words:"true"`
}

var Env Config

func LoadSettings() *Config {
	if err := godotenv.Load(); err != nil {
		slog.Warn(".env file not found, using default environment variables instead")
	}

	if err := envconfig.Process("datastar", &Env); err != nil {
		slog.Error("Failed to load environment configuration", "error", err)
		os.Exit(1)
	}

	return &Env
}
