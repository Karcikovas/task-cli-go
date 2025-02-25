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

func (c *Update) Run() {
	log.Println("CLI Update command doing stuff ")

}

func (c *Update) GetCmd() *console.Console {

	command := console.NewConsoleCommand("update", c.Run)

	return command
}
