package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"Stratum/internal/database"
	"Stratum/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r.Post("/register", handlers.RegisterHandler)
	r.Post("/login", handlers.LoginHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, (r)))
}
