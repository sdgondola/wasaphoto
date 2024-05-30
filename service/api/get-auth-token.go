package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/sdgondola/wasaphoto/service/database"
)

var ErrNoAuth = errors.New("unauthenticated")

func (rt *_router) getAuthToken(r *http.Request) (string, error) {
	token := strings.Split(r.Header.Get("Authorization"), "Bearer")[1]
	if token == "" {
		return token, ErrNoAuth
	}
	valid, err := rt.db.UserExists(token)
	if err != nil {
		return token, err
	}
	if !valid {
		return token, database.ErrUserNotFound
	}
	return token, nil
}
