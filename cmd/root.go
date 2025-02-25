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
	log.Println("Available cli command")

	for _, command := range r.commands {
		c := command.GetCmd()

		log.Println(fmt.Sprintf("task-cli: %s", c.Name))
	}
}
