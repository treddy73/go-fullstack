package route

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/treddy73/go-fullstack/internal/server/view"
)

func Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeHTML))

	r.Get("/", templ.Handler(view.Hello()).ServeHTTP)

	return r
}
