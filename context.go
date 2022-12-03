package durov

import "github.com/ayrat404/durov/client"

// RequestContext represents an incoming request from an user
type RequestContext struct {
	Command            string
	Params             map[string]string // key-value pairs parsed from RequestContext.InlineKeyboardData
	Text               string
	InlineKeyboardData string
	FileId             string
	ChatId             int
	Update             *client.Update
	TgClient           *client.TgClient
	UserInfo           UserInfo
}

// UserInfo contains user info
type UserInfo struct {
	UserName  string
	FirstName string
	LastName  string
}

func (r *RequestContext) SendMessage(msg string, inlineKeyboardParams map[string]string) (*client.Message, error) {
	params := &client.SendMessageParams{
		Text:   msg,
		ChatId: r.ChatId,
	}
	return r.TgClient.SendMessage(params)
}
