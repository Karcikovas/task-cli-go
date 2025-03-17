package task

import (
	"fmt"
	"regexp"
	"strconv"
	"task-cli-go/internal/console"
	"task-cli-go/internal/logger"
	"task-cli-go/internal/task"
)

type Update struct {
	task   task.Service
	logger logger.Service
}

func NewUpdate(task task.Service, logger logger.Service) *Update {
	return &Update{
		task:   task,
		logger: logger,
	}
}

func (c *Update) Run(args string) {
	idRegex := regexp.MustCompile(`\b\d+\b`)
	taskID := idRegex.FindString(args)
	descriptionRegex := regexp.MustCompile(`"([^"]+)"`)
	description := descriptionRegex.FindString(args)

	if len(taskID) == 0 || len(description) == 0 {
		c.logger.LogError("Wrong argument passed")
	}

	id, err := strconv.Atoi(taskID)

	if err != nil {
		c.logger.LogError(err.Error())
	}

	updated := c.task.UpdateTask(task.UpdateTaskDTO{
		ID:          id,
		Description: &description,
	})

	if updated {
		c.logger.LogSuccess(fmt.Sprintf(`Task %s updated`, taskID))
	} else {
		c.logger.LogWarning(fmt.Sprintf(`Failed to update %s`, taskID))
	}
}

func (c *Update) GetCmd() *console.Console {
	command := console.NewConsoleCommand("update", "update item based on item id", c.Run)

	return command
}
