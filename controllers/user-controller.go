package controllers

import (
	"net/http"

	"github.com/emcassi/open-stash-api/app"
	"github.com/emcassi/open-stash-api/models"
	"github.com/emcassi/open-stash-api/repos"
	"github.com/emcassi/open-stash-api/validation"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserWithEmailAndPassword(body models.UserCreationEmailPw) (*models.User, *app.AppError) {
	// Field Validation
	appErr := validation.ValidateUserName(body.Name)
	if appErr != nil {
		return nil, appErr
	}

	appErr = validation.ValidateUserEmail(body.Email)
	if appErr != nil {
		return nil, appErr
	}

	appErr = validation.ValidateUserPassword(body.Password)
	if appErr != nil {
		return nil, appErr
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, app.NewError(http.StatusInternalServerError, err)
	}

	body.Password = string(hash)

	user, appErr := repos.CreateUserWithEmailAndPassword(body)
	if appErr != nil {
		return nil, appErr
	}

	return user, nil
}
