package models

import (
	"time"

	"github.com/SilverLuhtoja/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DatabaseFeedFollowToFeedFollow(fedfol database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        fedfol.ID,
		UserID:    fedfol.UserID,
		FeedID:    fedfol.FeedID,
		CreatedAt: fedfol.CreatedAt,
		UpdatedAt: fedfol.UpdatedAt,
	}
}
