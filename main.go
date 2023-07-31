package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

type user struct {
	email string
	password string
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Servers
	fs := http.FileServer(http.Dir(""))
	mux := http.NewServeMux()

	// Handlers
	mux.Handle("/", fs)
	r := mux.new(route)
	http.ListenAndServe(":"+port, mux)

}

//dev-mux branch for mux development