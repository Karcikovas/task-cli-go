package task

import (
	"fmt"
	"strconv"
	"task-cli-go/internal/console"
	"task-cli-go/internal/logger"
	"task-cli-go/internal/task"
)

type List struct {
	service task.Service
	logger  logger.Service
}

func NewList(service task.Service, logger logger.Service) *List {
	return &List{
		service: service,
		logger:  logger,
	}
}

func (c *List) Run(_ string) {
	tasks := c.service.GetAllTasks()

	for _, t := range tasks {
		c.logger.LogInfo(fmt.Sprintf(`ID: %s Description: %s Updated: %s`, strconv.Itoa(*t.Id), t.Description, t.UpdatedAt))
	}
}

func (c *List) GetCmd() *console.Console {
	command := console.NewConsoleCommand("list", "list all items", c.Run)

	return command
}
