package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/huntermotko/gator/internal/database"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("username is required")
	}
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Println(user)
	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}
	fmt.Println("user has been set")
	return nil
}

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("register requires: [agg register username]")
	}
	ctx := context.Background()
	id := uuid.New()
	newUser := database.CreateUserParams{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}
	userRes, err := s.db.CreateUser(ctx, newUser)
	if err != nil {
		return err
	}
	if err := s.cfg.SetUser(userRes.Name); err != nil {
		return err
	}
	fmt.Println("User was created!")
	fmt.Println(userRes)

	return nil
}

func handlerCheckUsers(s *State, cmd Command) error {
	ctx := context.Background()
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return err
	}
	for _, u := range users {
		if u == s.cfg.Username {
			fmt.Printf("* %s (current)\n", u)
			continue
		}
		fmt.Printf("* %s\n", u)
	}

	return nil
}
