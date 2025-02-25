package main

import (
	"task-cli-go/cmd"
	taskCli "task-cli-go/cmd/cli/task"
	"task-cli-go/internal/console"
	//"task-cli-go/internal/storage"
	//"task-cli-go/internal/task"
)

func main() {

	//Register all cli commands
	create := taskCli.NewCreate()
	update := taskCli.NewUpdate()
	del := taskCli.NewDelete()

	//Init CLI Command and register them
	r := cmd.NewRoot([]console.Command{create, update, del})
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
