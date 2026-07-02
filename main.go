package main

import (
	"fmt"
	"gator/internal/config"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	programState := &state{cfg: &cfg}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("We require a command name")
		os.Exit(1)
	}
	err = cmds.run(programState, command{name: args[1], args: args[2:]})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
