package app

import (
	"fmt"
	"log/slog"
	"net/http"

	"aoc-runner/internal/config"
	"aoc-runner/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	Logger   *slog.Logger
	Router   *chi.Mux
	Settings *config.Settings
}

func New() *App {
	return &App{
		Settings: config.LoadSettings(),
	}
}

func (a *App) Initialize() error {
	a.Logger = config.NewColorLog(a.Settings.LogLevel)
	slog.SetDefault(a.Logger)

	a.Router = chi.NewRouter()
	a.Router.Use(middleware.Logger)

	// Make a new handler, pass what dependencies you want to initialize in the web handler
	// This way we can change out dependencies and keep all routes in /handler folder
	// May add other services later depending on need
	// Client Req --> Handler --> Client Res
	routeHandler := handler.NewHandler(a.Logger)
	routeHandler.SetupRoutes(a.Router)

	config.SetupFileServer(a.Logger, a.Router)

	return nil
}

func (a *App) Run() error {
	a.Logger.Info("Starting server", "port", a.Settings.Port)
	port := fmt.Sprintf(":%d", a.Settings.Port)
	return http.ListenAndServe(port, a.Router)
}
