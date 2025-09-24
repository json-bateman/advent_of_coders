package handler

import (
	"aoc-runner/views"
	"aoc-runner/views/stylesheet"
	"aoc-runner/views/y2018"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	l *slog.Logger
}

func NewHandler(l *slog.Logger) *Handler {
	return &Handler{
		l: l,
	}
}

func (h *Handler) SetupRoutes(r chi.Router) {
	r.Get("/", h.Index)
	r.Get("/stylesheet", h.Stylesheet)
	r.Get("/years", h.Years)
	r.Get("/year/2018", h.Y2018)

	// Solutions
	r.Post("/year/2018/sse", h.D01SSE)
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	component := views.IndexPage("Advent of Code")
	templ.Handler(component).ServeHTTP(w, r)
}

func (h *Handler) Stylesheet(w http.ResponseWriter, r *http.Request) {
	component := stylesheet.Stylesheet()
	templ.Handler(component).ServeHTTP(w, r)
}

func (h *Handler) Years(w http.ResponseWriter, r *http.Request) {
	component := views.Years()
	templ.Handler(component).ServeHTTP(w, r)
}

func (h *Handler) Y2018(w http.ResponseWriter, r *http.Request) {
	component := y2018.Y2018()
	templ.Handler(component).ServeHTTP(w, r)
}
