package task

import (
	"fmt"
	"log"
	"task-cli-go/internal/console"
	"task-cli-go/internal/task"
)

type Delete struct {
	service task.Service
}

func NewDelete(service task.Service) *Delete {
	return &Delete{service: service}
}

func (c *Delete) Run(args string) {
	deleted := c.service.DeleteTask(args)

	if deleted {
		log.Println(fmt.Sprintf(`Task %s deleted`, args))
	} else {
		log.Println("Unable to Delete Task")
	}
}

func (c *Delete) GetCmd() *console.Console {

	command := console.NewConsoleCommand("delete", "remove item from todo list by passing item id", c.Run)

	return command
}
