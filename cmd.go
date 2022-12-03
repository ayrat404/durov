package durov

type CommandDefinition struct {
	Name        string
	Description string
	Display     bool
}

type BotCommand interface {
	GetDefinition() CommandDefinition
	CanExecute(*RequestContext) bool
	Execute(*RequestContext)
}

type commandRouter struct {
	commands []BotCommand
}

func newRouter(commands []BotCommand) *commandRouter {
	return &commandRouter{
		commands: commands,
	}
}

func (c *commandRouter) Handle(ctx *RequestContext) {

}
