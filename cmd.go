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
	commands        []BotCommand
	cmdMap          map[string]BotCommand
	fallbackCommand BotCommand
}

func newRouter(commands []BotCommand, fallbackCommand BotCommand) *commandRouter {
	router := &commandRouter{
		commands:        commands,
		cmdMap:          make(map[string]BotCommand, len(commands)),
		fallbackCommand: fallbackCommand,
	}
	for _, command := range commands {
		router.cmdMap[command.GetDefinition().Name] = command
	}
	return router
}

func (c *commandRouter) Handle(request *RequestContext) {
	command := c.selectCommand(request)
	if command != nil {
		command.Execute(request)
	}
}

func (c *commandRouter) selectCommand(request *RequestContext) BotCommand {
	if request.Command != "" {
		cmd, ok := c.cmdMap[request.Command]
		if ok {
			return cmd
		}
	}
	for _, cmd := range c.commands {
		if cmd.CanExecute(request) {
			return cmd
		}
	}
	return c.fallbackCommand
}
