package task

import (
	"fmt"
	"log"
	"strconv"
	"task-cli-go/internal/console"
	"task-cli-go/internal/task"
)

type List struct {
	service task.Service
}

func NewList(service task.Service) *List {
	return &List{
		service: service,
	}
}

func (c *List) Run(_ string) {
	tasks := c.service.GetAllTasks()

	for _, t := range tasks {
		log.Println(fmt.Sprintf(`ID: %s Description: %s Updated: %s`, strconv.Itoa(*t.Id), t.Description, t.UpdatedAt))
	}
}

func (c *List) GetCmd() *console.Console {
	command := console.NewConsoleCommand("list", "list all items", c.Run)

	return command
}
