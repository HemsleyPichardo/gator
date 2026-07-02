package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("Login handler expects a username argument")
	}

	username := cmd.args[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("Logged in as", username)
	return nil
}
