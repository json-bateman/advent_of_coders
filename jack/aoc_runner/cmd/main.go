package main

import (
	"aoc"
	"aoc/web"
	"context"
	"fmt"
	"log"
	"log/slog"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	cfg := aoc.LoadSettings()

	slog.Info("~~~~Environment Vars~~~~~~~~~~")
	slog.Info("Development", "IsDev", cfg.IsDev)
	slog.Info("Port", "PORT", cfg.Port)
	slog.Info("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	if err := web.RunBlocking(ctx); err != nil {
		return fmt.Errorf("run web server: %w", err)
	}
	return nil
}
