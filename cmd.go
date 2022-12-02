package durov

type CommandDefinition struct {
	Name        string
	Description string
	Display     bool
}

type BotCommand interface {
	GetDefinition() CommandDefinition
	CanExecute(req *Request) bool
	Execute(req *Request)
}

type commandRouter struct {
	commands []BotCommand
}

func newRouter(commands []BotCommand) *commandRouter {
	return &commandRouter{
		commands: commands,
	}
}

func (c *commandRouter) Handle(req *Request) {

}
