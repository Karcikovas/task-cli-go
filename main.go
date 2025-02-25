package main

import (
	"os"
	"task-cli-go/cmd"
	"task-cli-go/cmd/Cli"
	//"task-Cli-go/internal/storage"
	//"task-Cli-go/internal/task"
)

func main() {
	cli := Cli.NewCLi()
	r := cmd.NewRoot(cli.GetCommands())

	args := os.Args
	// Check if the user provided an argument
	//if len(args) < 2 {
	//	fmt.Println("Usage: mycli <command>")
	//	return
	//}

	// Handle different commands
	command := args[1]

	switch command {
	case "create":
		cli.GetCommands()[0].Run()
	case "update":
		cli.GetCommands()[1].Run()
	case "delete":
		cli.GetCommands()[2].Run()
	default:
		r.AvailableCommands()
	}

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
