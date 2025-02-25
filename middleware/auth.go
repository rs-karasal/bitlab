package middleware

import (
	"encoding/base64"
	"fmt"
	"my_super_project/config"
	"my_super_project/database"
	"my_super_project/models"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func BasicAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Разбираем заголовок "Basic <base64-encoded-credentials>"
		authParts := strings.SplitN(authHeader, " ", 2)
		if len(authParts) != 2 || authParts[0] != "Basic" {
			http.Error(w, "Invalid Authorization format", http.StatusUnauthorized)
			return
		}

		decoded, err := base64.StdEncoding.DecodeString(authParts[1])
		if err != nil {
			http.Error(w, "Invalid Base64 encoding", http.StatusUnauthorized)
			return
		}

		credentials := strings.SplitN(string(decoded), ":", 2)
		if len(credentials) != 2 {
			http.Error(w, "Invalid credentials format", http.StatusUnauthorized)
			return
		}
		username, password := credentials[0], credentials[1]

		var user models.User
		err = database.Db.QueryRow("SELECT id, username, hashed_password FROM users WHERE username = $1", username).
			Scan(&user.ID, &user.Username, &user.HashedPassword)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if !user.CheckPassword(password) {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func JWTAuthMiddleware(next http.Handler, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpecred signing method: %v", token.Header["alg"])
			}
			return cfg.SecretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
