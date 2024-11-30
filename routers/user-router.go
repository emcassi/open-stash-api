package routers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/emcassi/open-stash-api/app"
	"github.com/emcassi/open-stash-api/controllers"
	"github.com/emcassi/open-stash-api/helpers"
	"github.com/emcassi/open-stash-api/models"
	"github.com/go-chi/chi/v5"
)

func HandleUserRoutes(r *chi.Mux) {
	r.Post("/users", createUser)
	r.Post("/auth/login", login)
}

func createUser(w http.ResponseWriter, r *http.Request) {

	method := r.URL.Query().Get("method")
	switch method {
	case "email_password":
		createUserEmailPassword(w, r)
	default:
		appErr := app.AppError{
			Status: http.StatusBadRequest,
			Error:  errors.New("invalid method provided"),
		}
		helpers.WriteError(w, appErr)
	}
}

func createUserEmailPassword(w http.ResponseWriter, r *http.Request) {
	var body models.UserCreationEmailPw

	bodyEmpty, err := IsRequestBodyEmpty(r)
	if err != nil {
		helpers.WriteError(w, app.AppError{Status: http.StatusBadRequest, Error: err})
		return
	}

	if bodyEmpty {
		helpers.WriteError(w, app.AppError{Status: http.StatusBadRequest, Error: errors.New("empty request body")})
		return
	}

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helpers.WriteError(w, app.AppError{Status: http.StatusBadRequest, Error: err})
		return
	}

	if body.Name == "" {
		helpers.WriteError(w, *app.NewError(http.StatusBadRequest, errors.New("field 'name' is required")))
		return
	}

	if body.Email == "" {
		helpers.WriteError(w, *app.NewError(http.StatusBadRequest, errors.New("field 'email' is required")))
		return
	}

	if body.Password == "" {
		helpers.WriteError(w, *app.NewError(http.StatusBadRequest, errors.New("field 'password' is required")))
		return
	}

	user, appErr := controllers.CreateUserWithEmailAndPassword(body)
	if appErr != nil {
		helpers.WriteError(w, *appErr)
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, map[string]interface{}{"message": fmt.Sprintf("Created User with Email: %s", *user.Email)})
}

func login(w http.ResponseWriter, r *http.Request) {
	method := r.URL.Query().Get("method")
	switch method {
	case "email_password":
		loginWithEmailAndPassword(w, r)
	default:
		appErr := app.AppError{
			Status: http.StatusBadRequest,
			Error:  errors.New("invalid method provided"),
		}
		helpers.WriteError(w, appErr)
	}
}

func loginWithEmailAndPassword(w http.ResponseWriter, r *http.Request) {
	var body models.UserLoginEmailPw	

	bodyEmpty, err := IsRequestBodyEmpty(r)
	if err != nil {
		helpers.WriteError(w, app.AppError{Status: http.StatusBadRequest, Error: err})
		return
	}

	if bodyEmpty {
		helpers.WriteError(w, app.AppError{Status: http.StatusBadRequest, Error: errors.New("empty request body")})
		return
	}

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helpers.WriteError(w, *app.NewError(http.StatusBadRequest, err))
		return
	}

	if body.Email == "" {
		helpers.WriteError(w, *app.NewError(http.StatusBadRequest, errors.New("field 'email' is required")))
		return
	}

	if body.Password == "" {
		helpers.WriteError(w, *app.NewError(http.StatusBadRequest, errors.New("field 'password' is required")))
		return
	}

	token, appErr := controllers.LoginWithEmailAndPassword(body)
	if appErr != nil {
		helpers.WriteError(w, *appErr)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, map[string]interface{}{"token": token})
}
