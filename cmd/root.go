package cmd

import (
	"fmt"
	"log"
	"task-cli-go/internal/console"
)

type Root struct {
	commands []console.Command
}

func NewRoot(commands []console.Command) *Root {
	return &Root{
		commands: commands,
	}
}

func (r *Root) Register() {
	log.Println("Available Cli command")

	for _, command := range r.commands {
		c := command.GetCmd()

		log.Println(fmt.Sprintf("task-Cli: %s", c.Name))
	}
}
