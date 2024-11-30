package validation

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/emcassi/open-stash-api/app"
)

func ValidateUserName(name string) *app.AppError {
	if len(name) == 0 {
		return app.NewError(http.StatusBadRequest, errors.New("field 'name' is required"))
	}

	if len(name) > 30 {
		return app.NewError(http.StatusBadRequest, errors.New("name must be no more than 30 characters long"))
	}

	return nil
}

func ValidateUserEmail(email string) *app.AppError {
	if len(email) == 0 {
		return app.NewError(http.StatusBadRequest, errors.New("field 'email' is required"))
	}

	verififer := emailverifier.NewVerifier()
	result, err := verififer.Verify(email)
	if err != nil {
		return app.NewError(http.StatusBadRequest, err)
	}

	if !result.Syntax.Valid {
		return app.NewError(http.StatusBadRequest, errors.New("email syntax is invalid"))
	}

	return nil
}

const PasswordSpecialChars = "!@#$%^&*-_=+,.<>/?\\|`~()[]{}"

func ValidateUserPassword(password string) *app.AppError {
	length := len(password)

	var issues []string

	if length == 0 {
		return app.NewError(http.StatusBadRequest, errors.New("field 'password' is required"))
	} else if length < 8 {
		issues = append(issues, "password must be at least 8 characters long")
	} else if length > 72 {
		issues = append(issues, "password must be no more than 72 characters long")
	}

	hasLower := false
	hasUpper := false
	hasDigit := false
	hasSpecial := false

	invalidChars := []rune{}

	for _, char := range password {
		switch {
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= '0' && char <= '9':
			hasDigit = true
		case strings.ContainsRune(PasswordSpecialChars, char):
			hasSpecial = true
		default:
			invalidChars = append(invalidChars, char)
		}
	}

	if !hasLower {
		issues = append(issues, "password must contain a lowercase letter")
	}

	if !hasUpper {
		issues = append(issues, "password must contain an uppercase letter")
	}

	if !hasDigit {
		issues = append(issues, "password must contain a number")
	}

	if !hasSpecial {
		issues = append(issues, fmt.Sprintf("password must contain at least one of the following special characters: `%s`", PasswordSpecialChars))
	}

	if len(invalidChars) == 1 {
		issues = append(issues, fmt.Sprintf("password contains invalid character: '%c'", invalidChars[0]))
	} else if len(invalidChars) > 1 {
		issues = append(issues, fmt.Sprintf("password contains invalid characters: '%s'", string(invalidChars)))
	}

	if len(issues) > 0 {
		return app.NewError(http.StatusBadRequest, errors.New(strings.Join(issues, "; ")))
	} else {
		return nil
	}
}
