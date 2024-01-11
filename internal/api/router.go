package api

import (
	"github.com/go-chi/chi"
)

func NewRouter(apiConfig *ApiConfig) *chi.Mux {
	api_router := NewApiCorsRouter()

	v1Router := chi.NewRouter()
	v1Router.Get("/readiness", ReadinessHandler)
	v1Router.Get("/err", ErrHandler)
	v1Router.Get("/users", apiConfig.GetUserByApiKey)
	v1Router.Post("/users", apiConfig.CreateUserHandler)

	api_router.Mount("/v1", v1Router)
	return api_router
}
