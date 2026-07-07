package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"strconv"
	"time"
)

func handlerBrowser(s *state, cmd command, user database.User) error {
	var limit int32
	if len(cmd.args) != 1 {
		// No limit given default to 2
		limit = 2
	} else {
		number, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %w", err)
		}
		limit = int32(number)
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: limit,
	})
	if err != nil {
		return fmt.Errorf("failed to get posts: %w", err)
	}
	for _, post := range posts {
		fmt.Printf("Title: %s\nURL: %s\nPublished: %v\n\n", post.Title, post.Url, post.PublishedAt.Time.Format(time.RFC1123))
	}
	return nil
}
