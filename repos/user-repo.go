package repos

import (
	"errors"
	"net/http"

	"github.com/emcassi/open-stash-api/app"
	"github.com/emcassi/open-stash-api/models"
	"gorm.io/gorm"
)

func CreateUserWithEmailAndPassword(body models.UserCreationEmailPw) (*models.User, *app.AppError) {
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

func GetUserByEmail(email string) (*models.User, *app.AppError) {
	var user models.User
	result := app.Db.Where("email = ?", email).First(&user)	
	if result.Error != nil {
		return nil, app.NewError(http.StatusBadRequest, result.Error)
	}

	return &user, nil
}
