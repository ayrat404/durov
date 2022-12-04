package durov

type BotParams struct {
	commands          []BotCommand
	middlewares       []func(Handler) Handler
	fallbackCommand   BotCommand
	customCmdExecutor Handler
}

type Handler func(*RequestContext)

func (b BotParams) Use(middlewares ...func(Handler) Handler) {
	b.middlewares = append(b.middlewares, middlewares...)
}

func (b BotParams) UseCmdExecutor(cmdExecutor Handler) {
	b.customCmdExecutor = cmdExecutor
}

func (b BotParams) AddCmd(commands ...BotCommand) {
	b.commands = append(b.commands, commands...)
}

func (b BotParams) AddFallbackCmd(command BotCommand) {
	b.fallbackCommand = command
}

func composeHandlers(middlewares []func(Handler) Handler, last Handler) Handler {
	if len(middlewares) == 0 {
		return last
	}
	handler := last
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
