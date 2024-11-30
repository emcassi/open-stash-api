package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/emcassi/open-stash-api/routers"
	"github.com/emcassi/open-stash-api/app"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	r := chi.NewRouter()

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
	
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
