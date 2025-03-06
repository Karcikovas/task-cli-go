package task

import (
	"fmt"
	"regexp"
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
	descriptionRegex := regexp.MustCompile(`"([^"]+)"`)
	description := descriptionRegex.FindString(args)

	if len(description) == 0 {
		c.logger.LogError("Wrong argument passed")

		return
	}

	saved, t := c.service.CreateTask(task.TaskDTO{
		Id:          nil,
		Description: args,
		Status:      "In Progress",
		CreatedAt:   nil,
		UpdatedAt:   nil,
	})

	if saved {
		c.logger.LogSuccess(fmt.Sprintf(`Task ID: %s`, strconv.Itoa(*t.Id)))
	} else {
		c.logger.LogWarning("FAILED to Add task")
	}
}

func (c *Add) GetCmd() *console.Console {

	command := console.NewConsoleCommand("add", "add new item into todo list", c.Run)

	return command
}
