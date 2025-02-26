package task

import (
	"log"
	"task-cli-go/internal/console"
)

type Update struct {
}

func NewUpdate() *Update {
	return &Update{}
}

func (c *Update) Run(args string) {
	log.Println("Cli Update command doing stuff ", args)

}

func (c *Update) GetCmd() *console.Console {

	command := console.NewConsoleCommand("update", "update item based on item id", c.Run)

	return command
}
