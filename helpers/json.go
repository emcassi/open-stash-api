package helpers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/emcassi/open-stash-api/app"
)

func WriteJSON(w http.ResponseWriter, status int, data map[string]interface{}) {
	json, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An error occurred. Plealse try again later"))
		return
	}

	w.WriteHeader(status)
	w.Write(json)
}

func WriteError(w http.ResponseWriter, appError app.AppError) {

	var data map[string]interface{}

	if appError.Status < 500 {
		data = map[string]interface{}{
			"error": appError.Error.Error(),
		}
	} else {
		data = map[string]interface{}{
			"error": "An error occurred. Please try again later.",
		}
		log.Printf("Error: %e\n", appError.Error)
	}

	json, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An error occurred. Plealse try again later"))
		return
	}

	w.WriteHeader(appError.Status)
	w.Write(json)
}
