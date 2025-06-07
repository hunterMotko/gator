package cmd

import (
	"context"
	"fmt"
)

func handlerReset(s *State, cmd Command) error {
	ctx := context.Background()
	if err := s.db.DeleteAllUsers(ctx); err != nil {
		return err
	}
	fmt.Println("Users Deleted")
	return nil
}
