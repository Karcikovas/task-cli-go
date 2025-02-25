package console

type Command interface {
	GetCmd() *Console
	Run()
}
