package y2025

import (
	"net/http"

	"github.com/a-h/templ"
)

func Y2025(w http.ResponseWriter, r *http.Request) {
	templ.Handler(y2025("AoC 2025")).Component.Render(r.Context(), w)
}
