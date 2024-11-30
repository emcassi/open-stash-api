package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OpenStash API"))
	})
}