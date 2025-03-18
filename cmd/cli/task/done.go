package task

import (
	"fmt"
	"regexp"
	"strconv"
	"task-cli-go/internal/console"
	"task-cli-go/internal/dto"
	"task-cli-go/internal/logger"
	"task-cli-go/internal/task"
)

type Done struct {
	task   task.Service
	logger logger.Service
}

func NewDone(task task.Service, logger logger.Service) *Done {
	return &Done{
		task:   task,
		logger: logger,
	}
}

func (c *Done) Run(args string) {
	idRegex := regexp.MustCompile(`\b\d+\b`)
	taskID := idRegex.FindString(args)

	if len(taskID) == 0 {
		c.logger.LogError("Wrong argument passed")
	}
	id, err := strconv.Atoi(taskID)

	if err != nil {
		c.logger.LogError(err.Error())
	}

	updated := c.task.UpdateTask(dto.UpdateTaskDTO{
		ID:     id,
		Status: &task.DONE,
	})

	if updated {
		c.logger.LogSuccess(fmt.Sprintf(`Task %s marked as done`, taskID))
	} else {
		c.logger.LogWarning(fmt.Sprintf(`Failed to mark task %s as done`, taskID))
	}
}

func (c *Done) GetCmd() *console.Console {
	command := console.NewConsoleCommand("mark-done", "mark task as done", c.Run)

	return command
}
