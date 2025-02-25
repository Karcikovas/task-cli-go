package console

type CommandAction func()

type Console struct {
	Name string
	run  CommandAction
}

func NewConsoleCommand(name string, runCommand CommandAction) *Console {
	return &Console{
		Name: name,
		run:  runCommand,
	}
}
