package main

import (
	"task-cli-go/cmd"
	"task-cli-go/cmd/Cli"
	//"task-Cli-go/internal/storage"
	//"task-Cli-go/internal/task"
)

func main() {

	cli := Cli.NewCLi()
	r := cmd.NewRoot(cli.GetCommands())
	r.Register()

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
