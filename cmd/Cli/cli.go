package Cli

import (
	taskCli "task-cli-go/cmd/Cli/task"
	"task-cli-go/internal/console"
	"task-cli-go/internal/logger"
	"task-cli-go/internal/task"
)

type Cli struct {
	commands map[string]console.Command
}

func NewCLi(
	task task.Service,
	logger logger.Service,
) *Cli {
	return &Cli{
		commands: map[string]console.Command{
			"add":    taskCli.NewAdd(task, logger),
			"update": taskCli.NewUpdate(task, logger),
			"delete": taskCli.NewDelete(task, logger),
			"list":   taskCli.NewList(task, logger),
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
