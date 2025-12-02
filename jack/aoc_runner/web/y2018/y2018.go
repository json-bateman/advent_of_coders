package y2018

import (
	"net/http"

	"github.com/a-h/templ"
)

func Y2018(w http.ResponseWriter, r *http.Request) {
	templ.Handler(y2018("AoC 2018")).Component.Render(r.Context(), w)
}
