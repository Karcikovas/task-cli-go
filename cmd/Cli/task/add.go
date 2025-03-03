package task

import (
	"math/rand"
	"task-cli-go/internal/console"
	"task-cli-go/internal/task"
	"time"
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

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) + 1

	c.task.CreateTask(task.TaskDTO{
		Id:          randomNumber,
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
