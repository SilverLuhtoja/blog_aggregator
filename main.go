package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SilverLuhtoja/blog_aggregator/internal/api"
	"github.com/SilverLuhtoja/blog_aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")
	PORT := os.Getenv("PORT")
	CONN := os.Getenv("CONN")

	db, err := sql.Open("postgres", CONN)
	if err != nil {
		log.Fatal("Couldn't connect to database")
	}
	dbQueries := database.New(db)
	apiConfig := &api.ApiConfig{DB: dbQueries, Client: &http.Client{}}

	r := api.NewRouter(apiConfig)
	server := &http.Server{
		Addr:        "localhost:" + PORT,
		Handler:     r,
		ReadTimeout: 5 * time.Second,
	}

	go func() {
		apiConfig.FeedWorker()
	}()

	log.Printf("Server running on: http://%s/v1\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
