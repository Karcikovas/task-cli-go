package task

import (
	"log"
	"os"
	"task-cli-go/internal/console"
)

type Create struct {
}

func NewCreate() *Create {
	return &Create{}
}

func (c *Create) Run() {
	args := os.Args

	log.Println("Create task ", args[1:])

	//TODO: Move this to internal cli command
	//External services which need to be removed
	//s := storage.CreateNewStorage()
	//t := task.CreateNewTask(s)
	//
	//t.CreateTask(task.TaskDTO{
	//	Id:          "1",
	//	Description: "Testing",
	//	Status:      "In Progress",
	//	CreatedAt:   nil,
	//	UpdatedAt:   nil,
	//})
}

func (c *Create) GetCmd() *console.Console {

	command := console.NewConsoleCommand("create", c.Run)

	return command
}
