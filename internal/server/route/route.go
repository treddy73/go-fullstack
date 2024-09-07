package route

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/treddy73/go-fullstack/internal/server/db"
	"github.com/treddy73/go-fullstack/internal/server/view"
	"net/http"
)

func Routes(todos *db.Collection) chi.Router {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeHTML))

	r.Get("/", templ.Handler(view.Hello(todos.Filter(""))).ServeHTTP)
	r.Post("/search", searchResultsHandler(todos))

	return r
}

func searchResultsHandler(todos *db.Collection) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()

		q := ""
		if r.Form.Has("q") {
			q = r.Form.Get("q")
		}

		component := view.SearchResults(todos.Filter(q))
		_ = component.Render(r.Context(), w)
	}
}
