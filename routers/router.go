package routers

import (
	"bytes"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OpenStash API"))
	})

	HandleUserRoutes(r)
}

func IsRequestBodyEmpty(r *http.Request) (bool, error) {
	buffer := make([]byte, 1)
	n, err := r.Body.Read(buffer)
	if err != nil && err != io.EOF {
		return false, err
	}
	if n == 0 {
		return true, nil
	}

	r.Body = io.NopCloser(io.MultiReader(bytes.NewReader(buffer[:n]), r.Body))
	return false, nil
}
