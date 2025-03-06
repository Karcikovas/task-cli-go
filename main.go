package main

import (
	"bufio"
	"os"
	"strings"
	"task-cli-go/cmd/Cli"
	cmdTask "task-cli-go/cmd/Cli/task"
	"task-cli-go/internal/logger"
	"task-cli-go/internal/storage"
	"task-cli-go/internal/task"
)

func main() {
	l := logger.NewLogger()
	s := storage.CreateNewStorage(l)
	t := task.CreateNewTask(s, l)
	cli := Cli.NewCLi(cmdTask.NewAdd(t, l))
	cli.SetAvailableCommands(
		cmdTask.NewAdd(t, l),
		cmdTask.NewDelete(t, l),
		cmdTask.NewDone(t, l),
		cmdTask.NewList(t, l),
		cmdTask.NewProgress(t, l),
		cmdTask.NewUpdate(t, l),
	)

	scanner := bufio.NewScanner(os.Stdin)
	cli.AvailableCommands()

	for {
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		c := parts[0]
		args := strings.Join(parts[1:], " ")

		command := cli.FindCommand(c)

		if command != nil {
			cli.SetCommand(command)
			cli.RunCommand(args)
		} else {
			l.LogError("Command not found")
			cli.AvailableCommands()
		}

	}
}
