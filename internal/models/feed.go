package models

import (
	"database/sql"
	"encoding/xml"
	"time"

	"github.com/SilverLuhtoja/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

// RSS struct represents the root structure of the RSS feed
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

// Channel struct represents the channel element in the RSS feed
type Channel struct {
	Title string `xml:"title"`
	// Link          string `xml:"link"`
	// Description   string `xml:"description"`
	// Generator     string `xml:"generator"`
	// Language      string `xml:"language"`
	// LastBuildDate string `xml:"lastBuildDate"`
	Items []Item `xml:"item"`
}

// Item struct represents the channel element in the RSS feed
type Item struct {
	Title         string `xml:"title"`
	Link          string `xml:"link"`
	PublishedDate string `xml:"pubDate"`
	Description   string `xml:"description"`
}

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
