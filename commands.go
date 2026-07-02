package main

import (
	"errors"
	"gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, handler func(*state, command) error) {
	c.registeredCommands[name] = handler
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.registeredCommands[cmd.name]
	if !ok {
		return errors.New("Command not found")
	}
	return handler(s, cmd)
}
