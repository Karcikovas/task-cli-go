package main

import (
	"bufio"
	"os"
	"strings"
	"task-cli-go/cmd"
	"task-cli-go/cmd/Cli"
	"task-cli-go/internal/logger"
	"task-cli-go/internal/storage"
	"task-cli-go/internal/task"
)

func main() {
	//TODO: move this like wire function for building dependencies
	l := logger.NewLogger()

	s := storage.CreateNewStorage(l)
	t := task.CreateNewTask(s, l)
	cli := Cli.NewCLi(t, l)
	r := cmd.NewRoot(cli.GetCommands())

	scanner := bufio.NewScanner(os.Stdin)

	for {
		r.AvailableCommands()

		scanner.Scan()
		input := scanner.Text()

		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		command := parts[0]
		args := strings.Join(parts[1:], " ")

		switch command {
		case "add":
			cli.CompleteCommand("add").Run(args)
		case "update":
			cli.CompleteCommand("update").Run(args)
		case "delete":
			cli.CompleteCommand("delete").Run(args)
		case "list":
			cli.CompleteCommand("list").Run(args)
		default:
			l.LogError("Command not found")
			r.AvailableCommands()
		}
	}
}
