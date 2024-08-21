package route

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/treddy73/go-fullstack/internal/server/view"
	"net/http"
	"strings"
)

func Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeHTML))

	r.Get("/", templ.Handler(view.Hello()).ServeHTTP)
	r.Post("/search", searchResultsHandler)

	return r
}

func searchResultsHandler(w http.ResponseWriter, r *http.Request) {
	db := []string{"one", "two", "three", "four", "five", "cat", "bird", "dog", "fish"}

	_ = r.ParseForm()

	var filter []string
	if r.Form.Has("q") {
		q := strings.ToLower(r.Form.Get("q"))
		for _, v := range db {
			if strings.Contains(v, q) {
				filter = append(filter, v)
			}
		}
	}

	component := view.SearchResults(filter)
	_ = component.Render(r.Context(), w)
}
