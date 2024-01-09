package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SilverLuhtoja/blog_aggregator/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	PORT := os.Getenv("PORT")
	r := api.NewRouter()

	server := &http.Server{
		Addr:        "localhost:" + PORT,
		Handler:     r,
		ReadTimeout: 5 * time.Second,
	}

	log.Printf("Server running on: http://%s/v1\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
