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

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
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

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("follow requires a feed url")
	}

	feedUrl := cmd.args[0]
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeedByURL(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	createdFeed, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Feed followed successfully:")
	fmt.Printf(" * FeedName:     %v\n", createdFeed.FeedName)
	fmt.Printf(" * UserName:     %v\n", createdFeed.UserName)

	return nil
}

func handlerListFeedFollows(s *state, cmd command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf(" * FeedName: %v, UserName: %v\n", feed.FeedName, feed.UserName)
	}

	return nil
}
