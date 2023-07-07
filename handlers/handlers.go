package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Handler struct {
	RedirectMap map[string]string
}

func (h *Handler) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (h *Handler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	endpoint := chi.URLParam(r, "endpoint")
	url, ok := h.RedirectMap[endpoint]
	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}
