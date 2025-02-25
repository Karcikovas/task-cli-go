package task

import (
	"log"
	"task-cli-go/internal/console"
)

type Delete struct {
}

func NewDelete() *Delete {
	return &Delete{}
}

func (c *Delete) Run() {
	log.Println("CLI Delete Command")
}

func (c *Delete) GetCmd() *console.Console {

	command := console.NewConsoleCommand("delete", c.Run)

	return command
}
