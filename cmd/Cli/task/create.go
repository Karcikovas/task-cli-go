package task

import (
	"log"
	"task-cli-go/internal/console"
)

type Create struct {
}

func NewCreate() *Create {
	return &Create{}
}

func (c *Create) Run() {
	log.Println("Cli Create Command")
}

func (c *Create) GetCmd() *console.Console {

	command := console.NewConsoleCommand("create", c.Run)

	return command
}
