package console

type CommandAction func(args string)

type Console struct {
	Name        string
	Description string
	run         CommandAction
}

func NewConsoleCommand(name string, description string, runCommand CommandAction) *Console {
	return &Console{
		Name:        name,
		Description: description,
		run:         runCommand,
	}
}
