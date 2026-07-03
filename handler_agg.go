package main

import (
	"context"
	"fmt"
)

func handlerAggregate(s *state, cmd command) error {
	ctx := context.Background()
	givenFeed := "https://www.wagslane.dev/index.xml"

	feed, err := fetchFeed(ctx, givenFeed)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", feed)

	return nil

}
