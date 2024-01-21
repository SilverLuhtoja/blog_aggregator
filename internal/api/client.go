package api

import (
	"context"
	"encoding/xml"
	"io"
	"time"

	"github.com/SilverLuhtoja/blog_aggregator/internal/database"
	"github.com/SilverLuhtoja/blog_aggregator/internal/models"
)

func (cfg *ApiConfig) ReturnModelDataFromUrl(r context.Context, feed database.Feed) (models.RSS, error) {

	res, err := cfg.Client.Get(feed.Url)
	if err != nil {
		return models.RSS{}, err
	}

	body, readErr := io.ReadAll(res.Body)
	defer res.Body.Close()
	if readErr != nil {
		return models.RSS{}, err
	}

	data := models.RSS{}
	jsonErr := xml.Unmarshal(body, &data)

	if jsonErr != nil {
		return models.RSS{}, err
	}

	err = cfg.DB.MarkFeedFetched(r, database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID:        feed.ID})
	if err != nil {
		return models.RSS{}, err
	}

	return data, nil
}
