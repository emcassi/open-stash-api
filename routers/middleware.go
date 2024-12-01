package routers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/emcassi/open-stash-api/app"
	"github.com/emcassi/open-stash-api/helpers"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			helpers.WriteError(w, *app.NewError(http.StatusInternalServerError, errors.New("environment variable 'JWT_SECRET' not found")))
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helpers.WriteError(w, *app.NewError(http.StatusUnauthorized, errors.New("authorization header missing")))
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			helpers.WriteError(w, *app.NewError(http.StatusUnauthorized, errors.New("invalid token format")))
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			helpers.WriteError(w, *app.NewError(http.StatusUnauthorized, errors.New("invalid or expired token")))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			expiration, err := time.Parse(time.RFC3339Nano, fmt.Sprint(claims["expiresAt"]))
			if err != nil {
				helpers.WriteError(w, *app.NewError(http.StatusUnauthorized, errors.New("invalid or expired token")))
				return
			}
			log.Println(expiration)
			if time.Now().Compare(expiration) > 0 {
				helpers.WriteError(w, *app.NewError(http.StatusUnauthorized, errors.New("invalid or expired token")))
				return
			}

			ctx := context.WithValue(r.Context(), app.TokenContextKey, claims)
			r = r.WithContext(ctx)
		} else {
			helpers.WriteError(w, *app.NewError(http.StatusUnauthorized, errors.New("invalid token claims")))
			return
		}

		next.ServeHTTP(w, r)
	})
}
