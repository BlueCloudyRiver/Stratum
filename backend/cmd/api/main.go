package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"Stratum/internal/database"
	"Stratum/internal/handlers"
	"Stratum/internal/handlers/projects"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

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

func main() {
	if err := godotenv.Load(".env", "../../../.env"); err != nil {
		log.Println("Warning: No .env file found, reading straight from system environment")
	}

	if err := database.Connect(); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	fmt.Println("Yay! Connected to PostgreSQL securely")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(corsMiddleware)
	r.Use(projectMiddleware)

	r.Post("/register", handlers.RegisterHandler)
	r.Post("/login", handlers.LoginHandler)
	r.Get("/projects", projects.GetProjectsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, (r)))
}
