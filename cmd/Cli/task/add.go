package task

import (
	"fmt"
	"log"
	"strconv"
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
	saved, t := c.task.CreateTask(task.TaskDTO{
		Id:          nil,
		Description: args,
		Status:      "In Progress",
		CreatedAt:   nil,
		UpdatedAt:   nil,
	})

	if saved {
		log.Println(fmt.Sprintf(`Task ID: %s`, strconv.Itoa(*t.Id)))
	} else {
		log.Println("FAILED to Add task")
	}

}

func (c *Add) GetCmd() *console.Console {

	command := console.NewConsoleCommand("add", "add new item into todo list", c.Run)

	return command
}
