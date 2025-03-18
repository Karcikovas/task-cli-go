package cli

import (
	"fmt"
	"task-cli-go/internal/console"
	"task-cli-go/internal/logger"
)

type Cli struct {
	command  console.Command
	commands []console.Command
	logger   logger.Service
}

func NewCLi(
	command console.Command,
	logger logger.Service,
) *Cli {
	return &Cli{
		command: command,
		logger:  logger,
	}
}

func (c *Cli) SetCommand(command console.Command) {
	c.command = command
}

func (c *Cli) RunCommand(args string) {
	c.command.Run(args)
}

func (c *Cli) SetAvailableCommands(commands ...console.Command) {
	c.commands = commands
}

func (c *Cli) AvailableCommands() {
	for _, command := range c.commands {
		cmd := command.GetCmd()

		c.logger.LogInfo(
			fmt.Sprintf(
				"task-cli: %s --- %s",
				cmd.Name,
				cmd.Description,
			))
	}
}

func (c *Cli) FindCommand(name string) console.Command {
	var command console.Command

	for _, v := range c.commands {
		cmd := v.GetCmd()

		if cmd.Name == name {
			command = v
		}
	}

	return command
}
