package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fs := http.FileServer(http.Dir(""))

	mux := http.NewServeMux()
	mux.Handle("/", fs)

	// mux.HandleFunc("/", indexHandler)

	http.ListenAndServe(":"+port, mux)

}
