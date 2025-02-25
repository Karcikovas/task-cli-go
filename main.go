package main

import (
	"os"
	"task-cli-go/cmd"
	"task-cli-go/cmd/Cli"
)

func main() {
	cli := Cli.NewCLi()
	r := cmd.NewRoot(cli.GetCommands())

	//TODO: move this into package
	args := os.Args
	if len(args) < 2 {
		r.AvailableCommands()
		return
	}

	command := args[1]

	//TODO: Need to remove this switch case to something more flexible
	switch command {
	case "create":
		cli.CompleteCommand("create").Run()
	case "update":
		cli.CompleteCommand("update").Run()
	case "delete":
		cli.CompleteCommand("delete").Run()
	default:
		r.AvailableCommands()
	}
}
