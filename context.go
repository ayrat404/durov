package durov

import "github.com/ayrat404/durov/client"

// RequestContext represents an incoming request from a user
type RequestContext struct {
	Command        string
	Text           string
	CallbackData   string
	FileId         string
	ChatId         int
	Update         *client.Update
	TgClient       *client.TgClient
	UserInfo       UserInfo
	InlineButtonId string
}

// UserInfo contains user info
type UserInfo struct {
	UserName  string
	FirstName string
	LastName  string
}

type InlineButton struct {
	Id    string
	Title string
}

func (r *RequestContext) Reply(msg string) (*client.Message, error) {
	return r.ReplyWithButtons(msg, nil)
}

func (r *RequestContext) ReplyWithButtons(msg string, inlineButtons []InlineButton) (*client.Message, error) {
	params := createMsgParams(r, msg, inlineButtons)
	return r.TgClient.SendMessage(params)
}

func createMsgParams(r *RequestContext, msg string, inlineButtons []InlineButton) *client.SendMessageParams {
	params := &client.SendMessageParams{
		Text:   msg,
		ChatId: r.ChatId,
	}
	if len(inlineButtons) > 0 {
		markup := client.InlineKeyboardMarkup{
			InlineKeyboard: make([]client.InlineKeyboardButton, 0, len(inlineButtons)),
		}
		for i, button := range inlineButtons {
			markup.InlineKeyboard[i] = client.InlineKeyboardButton{
				Text:         button.Title,
				CallbackData: "id=" + button.Id,
			}
		}
		params.ReplyMarkup = markup
	}
	return params
}
