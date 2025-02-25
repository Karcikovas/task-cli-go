package task

import (
	"log"
	"os"
	"task-cli-go/internal/console"
)

type Create struct {
}

func NewCreate() *Create {
	return &Create{}
}

func (c *Create) Run() {
	args := os.Args

	log.Println("Create task ", args[1:])
}

func (c *Create) GetCmd() *console.Console {

	command := console.NewConsoleCommand("create", c.Run)

	return command
}
