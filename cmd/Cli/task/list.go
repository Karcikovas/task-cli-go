package task

import (
	"log"
	"task-cli-go/internal/console"
	"task-cli-go/internal/task"
)

type List struct {
	task *task.Task
}

func NewList(task *task.Task) *List {
	return &List{
		task: task,
	}
}

func (c *List) Run(_ string) {
	tasks := c.task.GetAllTasks()

	for _, t := range tasks {
		log.Println(t.Description)
	}
}

func (c *List) GetCmd() *console.Console {
	command := console.NewConsoleCommand("list", "list all items", c.Run)

	return command
}
