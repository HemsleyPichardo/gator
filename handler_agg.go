package main

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerAggregate(s *state, cmd command) error {

	if len(cmd.args) != 1 {
		return fmt.Errorf("Expected time arugment")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Invalid time argument: %w", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			log.Println("Error scraping feeds:", err)
		}
	}

	return nil
}

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	next_feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(ctx, next_feed.ID)
	if err != nil {
		return err
	}

	feed, err := fetchFeed(ctx, next_feed.Url)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	layouts := []string{
		"Mon, 02 Jan 2006 15:04:05 -0700",
		"2006-01-02T15:04:05Z",
		// add more as needed
	}

	for _, item := range feed.Channel.Items {
		var publishedAt sql.NullTime
		for _, layout := range layouts {
			t, err := time.Parse(layout, item.PubDate)
			if err == nil {
				publishedAt = sql.NullTime{Time: t, Valid: true}
				break
			}
		} // if none matched, publishedAt.Valid remains false

		_, err := s.db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   now,
			UpdatedAt:   now,
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: publishedAt,
			FeedID:      next_feed.ID,
		})
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code == "23505" {
					continue
				}
			}
			log.Printf("error saving post: %v", err)
		}
	}
	return nil
}
