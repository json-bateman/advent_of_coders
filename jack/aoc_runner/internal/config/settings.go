package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const (
	PORT      = "PORT"
	LOG_LEVEL = "LOG_LEVEL"
)

type Settings struct {
	Port     int
	LogLevel slog.Level
}

func LoadSettings() *Settings {
	if err := godotenv.Load(); err != nil {
		slog.Warn(".env file not found, using environment variables instead")
	}

	config := &Settings{
		LogLevel: parseLogLevel(os.Getenv(LOG_LEVEL)),
		Port:     getEnvAsInt(PORT, 8080),
	}

	if err := config.validate(); err != nil {
		slog.Error("Configuration validation failed", "error", err)
		os.Exit(1)
	}

	config.logConfig()
	return config
}

func (c *Settings) validate() error {
	if c.Port < 1 || c.Port > 65535 {
		return fmt.Errorf("invalid port: %d", c.Port)
	}
	return nil
}

func (c *Settings) logConfig() {
	fmt.Println("------------------------------------------------------------------")
	slog.Info("Configuration loaded")
	slog.Info("PORT", "port", c.Port)
	slog.Info("LOG_LEVEL", "level", c.LogLevel)
	fmt.Println("------------------------------------------------------------------")
}

func getEnvAsString(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func parseLogLevel(level string) slog.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

