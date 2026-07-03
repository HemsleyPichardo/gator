package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("addfeed requires a name and url")
	}
	feedname := cmd.args[0]
	url := cmd.args[1]

	now := time.Now().UTC()
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      feedname,
		Url:       url,
		UserID:    currentUser.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Feed created successfully:")
	fmt.Printf(" * ID:         %v\n", feed.ID)
	fmt.Printf(" * CreatedAt:  %v\n", feed.CreatedAt)
	fmt.Printf(" * UpdatedAt:  %v\n", feed.UpdatedAt)
	fmt.Printf(" * Name:       %v\n", feed.Name)
	fmt.Printf(" * URL:        %v\n", feed.Url)
	fmt.Printf(" * UserID:     %v\n", feed.UserID)

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf(" * Name: %v, URL: %v, User: %v\n", feed.RssName, feed.Url, feed.Username)
	}

	return nil
}
