package Cli

import (
	taskCli "task-cli-go/cmd/Cli/task"
	"task-cli-go/internal/console"
	"task-cli-go/internal/task"
)

type Cli struct {
	commands map[string]console.Command
}

func NewCLi(task *task.Task) *Cli {
	return &Cli{
		commands: map[string]console.Command{
			"add":    taskCli.NewAdd(task),
			"update": taskCli.NewUpdate(),
			"delete": taskCli.NewDelete(task),
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

func (c *Cli) CompleteCommand(name string) console.Command {
	command := c.commands[name]

	return command
}
