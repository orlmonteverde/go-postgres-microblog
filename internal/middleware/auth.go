package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/orlmonteverde/go-postgres-microblog/pkg/claim"
	"github.com/orlmonteverde/go-postgres-microblog/pkg/response"
)

type key string

// Context keys
const (
	UserIDKey key = "id"
)

// Authorizator is a middleware that verifies if the token is valid.
func Authorizator(next http.Handler) http.Handler {
	signingString := os.Getenv("SIGNING_STRING")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		tokenString, err := tokenFromAuthorization(authorization)
		if err != nil {
			response.HTTPError(w, r, http.StatusUnauthorized, err.Error())
			return
		}

		c, err := claim.GetFromToken(tokenString, signingString)
		if err != nil {
			response.HTTPError(w, r, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserIDKey, c.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func tokenFromAuthorization(authorization string) (string, error) {
	if authorization == "" {
		return "", errors.New("autorization is required")
	}

	if !strings.HasPrefix(authorization, "Bearer") {
		return "", errors.New("invalid autorization format")
	}

	l := strings.Split(authorization, " ")
	if len(l) != 2 {
		return "", errors.New("invalid autorization format")
	}

	return l[1], nil
}
