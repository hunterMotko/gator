package cmd

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/huntermotko/gator/internal/database"
)

func handlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("URL required")
	}
	ctx := context.Background()
	feed, err := s.db.GetFeedByUrl(ctx, cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Println(feed)
	follows, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println(follows)
	return nil
}

func handlerFollowing(s *State, cmd Command, user database.User) error {
	ctx := context.Background()
	following, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return err
	}
	fmt.Println(following)
	for _, f := range following {
		fmt.Println(f.UserName, f.FeedName)
	}
	return nil
}

func handlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("URL required")
	}
	ctx := context.Background()
	err := s.db.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    cmd.Args[0],
	})
	if err != nil {
		return err
	}
	fmt.Printf("Deleted follow record\n")
	return nil
}
