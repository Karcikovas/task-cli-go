package task

import (
	"fmt"
	"log"
	"strconv"
	"task-cli-go/internal/console"
	"task-cli-go/internal/logger"
	"task-cli-go/internal/task"
)

type Add struct {
	service task.Service
	logger  logger.Service
}

func NewAdd(service task.Service, logger logger.Service) *Add {
	return &Add{
		service: service,
		logger:  logger,
	}
}

func (c *Add) Run(args string) {
	saved, t := c.service.CreateTask(task.TaskDTO{
		Id:          nil,
		Description: args,
		Status:      "In Progress",
		CreatedAt:   nil,
		UpdatedAt:   nil,
	})

	if saved {
		log.Println(fmt.Sprintf(`Task ID: %s`, strconv.Itoa(*t.Id)))
	} else {
		log.Println("FAILED to Add task")
	}

}

func (c *Add) GetCmd() *console.Console {

	command := console.NewConsoleCommand("add", "add new item into todo list", c.Run)

	return command
}
