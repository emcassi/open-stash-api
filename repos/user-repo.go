package repos

import (
	"errors"
	"log"
	"net/http"

	"github.com/emcassi/open-stash-api/app"
	"github.com/emcassi/open-stash-api/models"
	"gorm.io/gorm"
)

func CreateUserWithEmailAndPassword(body models.UserCreationEmailPw) (*models.User, *app.AppError) {
	log.Printf("BODY NAME: %s, EMAIL: %s, PASSWORD: %s\n", body.Name, body.Email, body.Password)
	user := models.User{Name: body.Name, Email: &body.Email, Password: &body.Password}
	result := app.Db.Create(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, app.NewError(http.StatusBadRequest, result.Error)
		} else {
			return nil, app.NewError(http.StatusInternalServerError, result.Error)
		}
	}

	return &user, nil
}
