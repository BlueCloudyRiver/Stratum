package middle

import (
	"Stratum/internal/handlers"
	"context"
	"encoding/json"
	"net/http"
	"os"

	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func projectMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "unauthorized",
			})
			return
		}
		if !strings.HasPrefix(authHeader, "Bearer") {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid request"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &handlers.Claims{}

		jwtSecret := []byte(os.Getenv("JWT_SECRET"))
		_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
			return jwtSecret, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
			return
		}

		userid := claims.UserID
		ctx := context.WithValue(r.Context(), "userid", userid)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
