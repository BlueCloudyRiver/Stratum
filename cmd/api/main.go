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

func main() {
	if err := godotenv.Load(".env", "../../.env"); err != nil {
		log.Println("Warning: No .env file found, reading straight from system environment")
	}

	if err := database.Connect(); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	fmt.Println("Yay! Connected to PostgreSQL securely")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/register", handlers.RegisterHandler)
	r.Post("/login", handlers.LoginHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
