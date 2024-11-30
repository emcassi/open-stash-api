package controllers

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/emcassi/open-stash-api/app"
	"github.com/emcassi/open-stash-api/models"
	"github.com/emcassi/open-stash-api/repos"
	"github.com/emcassi/open-stash-api/validation"
	"github.com/golang-jwt/jwt/v5"
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

func LoginWithEmailAndPassword(body models.UserLoginEmailPw) (string, *app.AppError) {
	user, appErr := repos.GetUserByEmail(body.Email)	
	if appErr != nil {
		return "", appErr
	}

	err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(body.Password))
	if err != nil {
		return "", app.NewError(http.StatusBadRequest, errors.New("password is incorrect"))
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", app.NewError(http.StatusInternalServerError, errors.New("environment variable 'JWT_SECRET' not found"))
	}

	expiration := time.Now().Add(time.Hour * 24 * 30)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"expiresAt": expiration,
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", app.NewError(http.StatusInternalServerError, err)
	}

	return tokenString, nil
}
