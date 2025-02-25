package Cli

import (
	taskCli "task-cli-go/cmd/Cli/task"
	"task-cli-go/internal/console"
)

type Cli struct {
	commands map[string]console.Command
}

func NewCLi() *Cli {
	return &Cli{
		commands: map[string]console.Command{
			"create": taskCli.NewCreate(),
			"update": taskCli.NewUpdate(),
			"delete": taskCli.NewDelete(),
		},
	}
}

func (c *Cli) GetCommands() []console.Command {
	var commands = make([]console.Command, 0)

	for _, command := range c.commands {
		commands = append(commands, command)
	}

	return commands
}
