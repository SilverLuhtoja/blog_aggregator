package api

import (
	"github.com/SilverLuhtoja/blog_aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type ApiConfig struct {
	DB *database.Queries
}

func NewApiCorsRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:         300,
	}))
	return router
}
