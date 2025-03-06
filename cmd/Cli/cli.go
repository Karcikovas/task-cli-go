package Cli

import (
	"fmt"
	"log"
	"task-cli-go/internal/console"
)

type Cli struct {
	command  console.Command
	commands []console.Command
}

func NewCLi(command console.Command) *Cli {
	return &Cli{
		command: command,
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
		c := command.GetCmd()

		log.Println(fmt.Sprintf("task-Cli: %s --- %s", c.Name, c.Description))
	}
}

func (c *Cli) FindCommand(name string) console.Command {
	var command console.Command = nil

	for _, v := range c.commands {
		cmd := v.GetCmd()

		if cmd.Name == name {
			command = v
		}
	}

	return command
}
