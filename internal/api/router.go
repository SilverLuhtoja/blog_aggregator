package api

import "github.com/go-chi/chi"

func NewRouter() *chi.Mux {
	api_router := NewApiCorsRouter()
	router := chi.NewRouter()

	api_router.Mount("/v1", router)

	return api_router
}
