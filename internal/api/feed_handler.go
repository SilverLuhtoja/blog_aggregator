package api

import (
	"encoding/xml"
	"fmt"
	"io"
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

// RSS struct represents the root structure of the RSS feed
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

// Channel struct represents the channel element in the RSS feed
type Channel struct {
	Title         string `xml:"title"`
	Link          string `xml:"link"`
	Description   string `xml:"description"`
	Generator     string `xml:"generator"`
	Language      string `xml:"language"`
	LastBuildDate string `xml:"lastBuildDate"`
}

func (cfg *ApiConfig) FetchFeedData(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := cfg.DB.GetNextFeedsToFetch(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	client := http.Client{}

	feed := feeds[0]
	res, err := client.Get(feed.Url)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	body, readErr := io.ReadAll(res.Body)
	defer res.Body.Close()
	if readErr != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := RSS{}
	jsonErr := xml.Unmarshal(body, &data)
	if jsonErr != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = cfg.DB.MarkFeedFetched(r.Context(), database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID:        feed.ID})

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, 200, data)
}
