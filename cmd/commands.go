package cmd

import (
	"fmt"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Cmds map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	fn, ok := c.Cmds[cmd.Name]
	if !ok {
		return fmt.Errorf("Not a command")
	}
	if err := fn(s, cmd); err != nil {
		return err
	}
	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Cmds[name] = f
}
