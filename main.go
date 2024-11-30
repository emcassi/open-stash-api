package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/emcassi/open-stash-api/app"
	"github.com/emcassi/open-stash-api/routers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	r := chi.NewRouter()

  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)

	envFile := "dev.env"
	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf(".env file not found: %s\n", envFile)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = app.DefaultPort
	}

	routers.HandleRoutes(r)
	
	log.Printf("Listening at: http://127.0.0.1:%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
