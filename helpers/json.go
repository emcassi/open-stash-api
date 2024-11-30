package helpers

import (
	"encoding/json"
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
	data := map[string]interface{}{
		"message": appError.Message,
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
