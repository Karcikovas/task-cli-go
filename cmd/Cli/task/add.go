package task

import (
	"task-cli-go/internal/console"
	"task-cli-go/internal/task"
)

type Add struct {
	task *task.Task
}

func NewAdd(task *task.Task) *Add {
	return &Add{
		task: task,
	}
}

func (c *Add) Run(args string) {
	c.task.CreateTask(task.TaskDTO{
		Id:          nil,
		Description: args,
		Status:      "In Progress",
		CreatedAt:   nil,
		UpdatedAt:   nil,
	})
}

func (c *Add) GetCmd() *console.Console {

	command := console.NewConsoleCommand("add", "add new item into todo list", c.Run)

	return command
}
