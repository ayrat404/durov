package durov

type CommandDefinition struct {
	Name        string
	Description string
	Display     bool
}

type BotCommand interface {
	GetDefinition() CommandDefinition
	Execute(*RequestContext)
}

type CommandExecutor struct {
	commands        []BotCommand
	cmdMap          map[string]BotCommand
	fallbackCommand BotCommand
}

func newCommandExecutor(commands []BotCommand, fallbackCommand BotCommand) *CommandExecutor {
	router := &CommandExecutor{
		commands:        commands,
		cmdMap:          make(map[string]BotCommand, len(commands)),
		fallbackCommand: fallbackCommand,
	}
	for _, command := range commands {
		router.cmdMap[command.GetDefinition().Name] = command
	}
	return router
}

func (c *CommandExecutor) Execute(request *RequestContext) {
	command := c.selectCommand(request)
	if command != nil {
		command.Execute(request)
	}
}

func (c *CommandExecutor) selectCommand(request *RequestContext) BotCommand {
	if request.Command != "" {
		cmd, ok := c.cmdMap[request.Command]
		if ok {
			return cmd
		}
	}
	return c.fallbackCommand
}
