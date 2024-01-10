package api

import "github.com/go-chi/chi"

func NewRouter() *chi.Mux {
	api_router := NewApiCorsRouter()

	v1Router := chi.NewRouter()
	v1Router.Get("/readiness", ReadinessHandler)
	v1Router.Get("/err", ErrHandler)

	api_router.Mount("/v1", v1Router)
	return api_router
}
