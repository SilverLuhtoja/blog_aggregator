package api

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/SilverLuhtoja/blog_aggregator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (cfg *ApiConfig) FeedWorker() {
	var wg sync.WaitGroup
	for {
		ctx := context.TODO()
		feeds, err := cfg.DB.GetNextFeedsToFetch(ctx)
		if err != nil {
			log.Fatal(err)
		}
		for _, feed := range feeds {
			// Increment the WaitGroup counter.
			wg.Add(1)
			// Launch a goroutine to fetch the URL.
			go func(feed database.Feed) {
				// Decrement the counter when the goroutine completes.
				defer wg.Done()
				// Fetch the URL.
				res, err := cfg.ReturnModelDataFromUrl(ctx, feed)
				if err != nil {
					log.Fatal(err)
				}
				for _, item := range res.Channel.Items {
					_, err = cfg.DB.CreatePost(ctx, database.CreatePostParams{
						ID:          uuid.New(),
						CreatedAt:   time.Now().UTC(),
						UpdatedAt:   time.Now().UTC(),
						Title:       item.Title,
						Url:         item.Link,
						PublishedAt: *parseDate(item.PublishedDate),
						Description: item.Description,
						FeedID:      feed.ID,
					})
					if err != nil {
						if !isDuplicateKeyError(err) {
							fmt.Println("CreatePost error:", err)
							fmt.Println(err)
							return
						}
					}
				}
			}(feed)
		}

		// Wait for all HTTP fetches to complete.
		wg.Wait()
		fmt.Println("FETCHED")
		time.Sleep(10 * time.Second)
	}
}

func parseDate(inputTime string) *time.Time {
	inputLayout := "Mon, 02 Jan 2006 15:04:05 -0700"

	// Parse the input time string into a time.Time value
	parsedTime, err := time.Parse(inputLayout, inputTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil
	}
	return &parsedTime
}

func isDuplicateKeyError(err error) bool {
	pqError, ok := err.(*pq.Error)
	return ok && pqError.Code == "23505" // PostgreSQL error code for unique violation
}
