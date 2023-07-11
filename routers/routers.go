package routers

import (
	"net/http"

	"github.org/joshuabl97/chichichi/handlers"

	"github.com/go-chi/chi"
)

func NewRouter(handler *handlers.Handler) http.Handler {
	r := chi.NewRouter()

	r.Get("/", handler.HelloWorldHandler)
	r.Get("/{endpoint}", handler.RedirectHandler)

	return r
}
