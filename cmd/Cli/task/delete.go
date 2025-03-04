package task

import (
	"log"
	"task-cli-go/internal/console"
	"task-cli-go/internal/task"
)

type Delete struct {
	task *task.Task
}

func NewDelete(task *task.Task) *Delete {
	return &Delete{task: task}
}

func (c *Delete) Run(args string) {
	c.task.DeleteTask("0")

	log.Println(`Task Deleted`)
}

func (c *Delete) GetCmd() *console.Console {

	command := console.NewConsoleCommand("delete", "remove item from todo list by passing item id", c.Run)

	return command
}
