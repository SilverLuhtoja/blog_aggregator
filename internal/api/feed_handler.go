package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SilverLuhtoja/blog_aggregator/internal/database"
	"github.com/SilverLuhtoja/blog_aggregator/internal/models"
	"github.com/google/uuid"
)

type Feed struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (cfg *ApiConfig) CreateFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	params, err := GetParamsFromRequestBody(Feed{}, r)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("createFeedHandler - ", err))
		return
	}
	feed_id := uuid.New()

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        feed_id,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("createFeedHandler - ", err))
		return
	}

	feed_follow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed_id,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("createFeedHandler - ", err))
		return
	}

	type response struct {
		Feed       models.Feed       `json:"feed"`
		FeedFollow models.FeedFollow `json:"feed_follow"`
	}

	RespondWithJSON(w, http.StatusCreated, response{Feed: models.DatabaseFeedToFeed(feed), FeedFollow: models.DatabaseFeedFollowToFeedFollow(feed_follow)})
}

func (cfg *ApiConfig) GetAllFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("getAllFeedsHandler - ", err))
		return
	}

	RespondWithJSON(w, http.StatusCreated, feeds)
}

func (cfg *ApiConfig) FetchFeedData(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := cfg.DB.GetNextFeedsToFetch(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data, err := cfg.ReturnModelDataFromUrl(r.Context(), feeds[0])
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, 200, data)
}
