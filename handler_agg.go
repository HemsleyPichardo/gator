package main

import (
	"context"
	"fmt"
	"time"
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
		scrapeFeeds(s)
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

	for _, item := range feed.Channel.Items {
		fmt.Println(item.Title)
	}

	return nil
}
