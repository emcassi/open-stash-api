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

func CreateUserWithEmailAndPassword(body models.UserCreationEmailPw) (string, *app.AppError) {
	// Field Validation
	appErr := validation.ValidateUserName(body.Name)
	if appErr != nil {
		return "", appErr
	}

	appErr = validation.ValidateUserEmail(body.Email)
	if appErr != nil {
		return "", appErr
	}

	appErr = validation.ValidateUserPassword(body.Password)
	if appErr != nil {
		return "", appErr
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", app.NewError(http.StatusInternalServerError, err)
	}

	body.Password = string(hash)

	user, appErr := repos.CreateUserWithEmailAndPassword(body)
	if appErr != nil {
		return "", appErr
	}

	tokenString, appErr := CreateJWT(user.ID)
	if appErr != nil {
		return "", appErr
	}

	return tokenString, nil
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

	tokenString, appErr := CreateJWT(user.ID)
	if appErr != nil {
		return "", appErr
	}

	return tokenString, nil
}

func CreateJWT(userId string) (string, *app.AppError) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", app.NewError(http.StatusInternalServerError, errors.New("environment variable 'JWT_SECRET' not found"))
	}

	expiration := time.Now().Add(time.Hour * 24 * 30)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userId,
		"expiresAt": expiration,
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", app.NewError(http.StatusInternalServerError, err)
	}

	return tokenString, nil
}
