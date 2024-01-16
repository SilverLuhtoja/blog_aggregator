package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/SilverLuhtoja/blog_aggregator/internal/database"
	"github.com/SilverLuhtoja/blog_aggregator/internal/models"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type FeedFollow struct {
	FeedId uuid.UUID `json:"feed_id"`
}

func (cfg *ApiConfig) CreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	params, err := GetParamsFromRequestBody(FeedFollow{}, r)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("createFeedFollow - ", err))
		return
	}

	RespondWithJSON(w, http.StatusCreated, models.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (cfg *ApiConfig) DeleteFeedFollowById(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "feedFollowID")
	id, err := uuid.Parse(params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = cfg.doesValidFeedFollowExist(r.Context(), id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	message := fmt.Sprintf("FeedFollow with id= %s deleted", id)
	RespondWithJSON(w, 200, message)
}

func (cfg *ApiConfig) GetAllUserFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := cfg.DB.GetAllUserFeedFollows(r.Context(), user.ID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	feed_follows := []models.FeedFollow{}
	for _, follow := range feedFollows {
		feed_follows = append(feed_follows, models.DatabaseFeedFollowToFeedFollow(follow))
	}
	RespondWithJSON(w, 200, feed_follows)
}

func (cfg *ApiConfig) doesValidFeedFollowExist(ctx context.Context, id uuid.UUID) error {
	if _, ok := cfg.DB.FindFeedFollowById(ctx, id); ok != nil {
		return errors.New("feed follow does not exist")
	}
	return nil
}
