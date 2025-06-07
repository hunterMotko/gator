package cmd

import (
	"context"

	"github.com/huntermotko/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		ctx := context.Background()
		user, err := s.db.GetUser(ctx, s.cfg.Username)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
