package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	service "github.com/mktsy/go-webhook/controllers"
	lib "github.com/mktsy/go-webhook/shared"
)

func init() {
	// load config file
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file! Try get a path...")
		if err2 := godotenv.Load(lib.GetPath() + "/.env"); err2 != nil {
			log.Printf("Fail...")
			os.Exit(1)
		}
	}
}

func main() {
	r := chi.NewRouter()
	r.HandleFunc("/", service.HandlerMessenger)
	port := os.Getenv("PORT")
	log.Printf("Server started on localhost%s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
