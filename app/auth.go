package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/kerrai1990/api_contacts/models"
	u "github.com/kerrai1990/api_contacts/utils"
)

// JwtAuthentication -
var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Urls that don't need authentication
		notAuth := []string{
			"/api/users",
			"/api/session",
		}

		// Get the current request path
		requestPath := r.URL.Path

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
			}
		}

		response := make(map[string]interface{})

		// Get Authorization header
		tokenHeader := r.Header.Get("Authorization")

		// Validate token
		if tokenHeader == "" {
			response = u.Message(false, "Missing Auth Token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		// Split token from Barer in string
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		givenToken := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(givenToken, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_PASSWORD")), nil
		})

		// Malformed token (403)
		if err != nil {
			response = u.Message(false, "Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		// Invalid Token/Not Found - checked on line 53 (ParseWithClaims)
		if !token.Valid {
			response = u.Message(false, "Token is not valid")
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		fmt.Sprintf("User: %s", tk.UserID)
		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}
