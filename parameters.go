package durov

import "github.com/ayrat404/durov/client"

// Request represents an incoming request from an user
type Request struct {
	Command  string
	Params   map[string]string
	Text     string
	File     []byte
	ChatId   int
	UserInfo UserInfo
	Update   *client.Update
	TgClient *client.TgClient
}

type Response struct {
}

// UserInfo contains user info
type UserInfo struct {
	UserName  string
	FirstName string
	LastName  string
}

type BotParams struct {
	commands    []BotCommand
	middlewares []func(Handler) Handler
}

func (b BotParams) Use(middlewares ...func(Handler) Handler) {
	b.middlewares = append(b.middlewares, middlewares...)
}

func (b BotParams) AddCmd(commands ...BotCommand) {
	b.commands = append(b.commands, commands...)
}
