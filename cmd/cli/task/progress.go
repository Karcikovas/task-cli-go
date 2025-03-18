package task

import (
	"fmt"
	"regexp"
	"strconv"
	"task-cli-go/internal/console"
	"task-cli-go/internal/logger"
	"task-cli-go/internal/task"
)

type Progress struct {
	task   task.Service
	logger logger.Service
}

func NewProgress(task task.Service, logger logger.Service) *Progress {
	return &Progress{
		task:   task,
		logger: logger,
	}
}

func (c *Progress) Run(args string) {
	idRegex := regexp.MustCompile(`\b\d+\b`)
	taskID := idRegex.FindString(args)

	if len(taskID) == 0 {
		c.logger.LogError("Wrong argument passed")
	}
	id, err := strconv.Atoi(taskID)

	if err != nil {
		c.logger.LogError(err.Error())
	}

	updated := c.task.UpdateTask(task.UpdateTaskDTO{
		ID:     id,
		Status: &task.InPROGRESS,
	})

	if updated {
		c.logger.LogSuccess(fmt.Sprintf(`Task %s marked as InProgress`, taskID))
	} else {
		c.logger.LogWarning(fmt.Sprintf(`Failed to mark task %s as InProgress`, taskID))
	}
}

func (c *Progress) GetCmd() *console.Console {
	command := console.NewConsoleCommand("in-progress", "mark task as in progress", c.Run)

	return command
}
