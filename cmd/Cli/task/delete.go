package task

import (
	"fmt"
	"regexp"
	"task-cli-go/internal/console"
	"task-cli-go/internal/logger"
	"task-cli-go/internal/task"
)

type Delete struct {
	service task.Service
	logger  logger.Service
}

func NewDelete(service task.Service, logger logger.Service) *Delete {
	return &Delete{
		service: service,
		logger:  logger,
	}
}

func (c *Delete) Run(args string) {
	idRegex := regexp.MustCompile(`\b\d+\b`)
	taskID := idRegex.FindString(args)

	if len(taskID) == 0 {
		c.logger.LogError("Wrong argument passed")
	}

	deleted := c.service.DeleteTask(args)

	if deleted {
		c.logger.LogSuccess(fmt.Sprintf(`Task %s deleted`, args))
	} else {
		c.logger.LogWarning("Unable to Delete Task")
	}
}

func (c *Delete) GetCmd() *console.Console {

	command := console.NewConsoleCommand("delete", "remove item from todo list by passing item id", c.Run)

	return command
}
