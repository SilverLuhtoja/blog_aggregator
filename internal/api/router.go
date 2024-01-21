package api

import (
	"github.com/go-chi/chi"
)

func NewRouter(cfg *ApiConfig) *chi.Mux {
	api_router := NewApiCorsRouter()

	v1Router := chi.NewRouter()
	v1Router.Get("/readiness", ReadinessHandler)
	v1Router.Get("/err", ErrHandler)
	v1Router.Get("/users", cfg.middlewareAuth(cfg.GetUserHandler))
	v1Router.Post("/users", cfg.CreateUserHandler)

	v1Router.Get("/feeds", cfg.GetAllFeedsHandler)
	v1Router.Post("/feeds", cfg.middlewareAuth(cfg.CreateFeedHandler))
	v1Router.Get("/feeds/fetch", cfg.middlewareAuth(cfg.FetchFeedData))

	v1Router.Get("/feed_follows", cfg.middlewareAuth(cfg.GetAllUserFeedFollows))
	v1Router.Post("/feed_follows", cfg.middlewareAuth(cfg.CreateFeedFollow))
	v1Router.Delete("/feed_follows/{feedFollowID}", cfg.DeleteFeedFollowById)

	v1Router.Get("/posts", cfg.middlewareAuth(cfg.GetPostsByUser))

	api_router.Mount("/v1", v1Router)
	return api_router
}
