package main

import (
	"bufio"
	"os"
	"strings"
	"task-cli-go/cmd/Cli"
	"task-cli-go/internal/logger"
)

type Application struct {
	logger logger.Service
	cli    Cli.Cli
}

func NewApplication(logger logger.Service, cli Cli.Cli) *Application {
	return &Application{
		logger,
		cli,
	}
}

func (a *Application) Start() error {
	scanner := bufio.NewScanner(os.Stdin)
	a.cli.AvailableCommands()

	for {
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		c := parts[0]
		args := strings.Join(parts[1:], " ")
		command := a.cli.FindCommand(c)

		if command != nil {
			a.cli.SetCommand(command)
			a.cli.RunCommand(args)
		} else {
			a.logger.LogError("Command not found")
			a.cli.AvailableCommands()
		}
	}
}
