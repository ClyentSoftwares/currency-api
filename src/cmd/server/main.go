package main

import (
	"log"
	"net/http"
	"os"

	"github.com/clyentsoftwares/currency-api/src/internal/app/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	server := server.NewServer()
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, server))
}
