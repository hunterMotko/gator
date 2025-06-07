package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/huntermotko/gator/internal/database"
)

func handlerAgg(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			return err
		}
	}
}

func handlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("name and url required")
	}
	ctx := context.Background()
	feedRecord, err := s.db.AddFeed(ctx, database.AddFeedParams{
		ID:     uuid.New(),
		Name:   cmd.Args[0],
		Url:    cmd.Args[1],
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	ff, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: feedRecord.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println("feed follow", ff)
	return nil
}

func handlerFeeds(s *State, cmd Command) error {
	ctx := context.Background()
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Printf("%s - %s - %s\n", feed.FeedName, feed.Url, feed.UserName)
	}
	return nil
}
