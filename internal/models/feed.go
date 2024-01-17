package models

import (
	"database/sql"
	"time"

	"github.com/SilverLuhtoja/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

type Feed struct {
	ID            uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	Url           string    `json:"url"`
	UserID        uuid.UUID `json:"user_id"`
	LastFetchedAt time.Time `json:"last_fetched_at"`
}

func DatabaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:            feed.ID,
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.UpdatedAt,
		Name:          feed.Name,
		Url:           feed.Url,
		UserID:        feed.UserID,
		LastFetchedAt: convertNullTime(feed.LastFetchedAt),
	}
}

// Function to convert sql.NullTime to time.Time
func convertNullTime(nullTime sql.NullTime) time.Time {
	if nullTime.Valid {
		return nullTime.Time
	}
	// You can return a default value or handle the NULL case as needed
	return time.Time{}
}
