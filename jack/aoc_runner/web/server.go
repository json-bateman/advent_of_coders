package web

import (
	"aoc"
	"aoc/web/components"
	"aoc/web/y2018"
	"aoc/web/y2025"
	"context"
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/benbjohnson/hashfs"
	"github.com/go-chi/chi/v5"
)

//go:embed static/*
var staticFS embed.FS

func init() {
	// Initialize the static file system for components to use
	components.StaticSys = hashfs.NewFS(staticFS)
}

func setupRoutes() chi.Router {
	r := chi.NewRouter()

	r.Handle("/static/*", hashfs.FileServer(components.StaticSys))
	r.Get("/", home)

	r.Get("/year", yearsHandler)
	// 2018
	r.Get("/year/2018", y2018.Y2018)
	r.Get("/year/2018/day/1/part/1", y2018.D1P1SSE)
	r.Get("/year/2018/day/1/part/2", y2018.D1P2SSE)
	// 2025
	r.Get("/year/2025", y2025.Y2025)
	r.Get("/year/2025/day/1/part/1", y2025.D1P1SSE)
	r.Get("/year/2025/day/1/part/2", y2025.D1P2SSE)
	return r
}

func home(w http.ResponseWriter, r *http.Request) {
	index := templ.Handler(index("Advent of Coders"))
	index.Component.Render(r.Context(), w)
}

func yearsHandler(w http.ResponseWriter, r *http.Request) {
	years := templ.Handler(years("AoC Years"))
	years.Component.Render(r.Context(), w)
}

// RunBlocking sets up routes, starts the server, handles cleanup
func RunBlocking(setupCtx context.Context) error {
	router := setupRoutes()

	addr := fmt.Sprintf(":%d", aoc.Env.Port)
	srv := http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		<-setupCtx.Done()
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down server: %v", err)
		}
	}()

	log.Printf("Starting server on http://localhost%s", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Printf("Error starting server: %v", err)
	}
	return nil
}
