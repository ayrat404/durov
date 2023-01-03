package durov

import (
	"github.com/ayrat404/durov/client"
	"net/url"
)

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
	Commands       []BotCommand
}

// UserInfo contains user info
type UserInfo struct {
	UserName  string
	FirstName string
	LastName  string
}

type InlineButton struct {
	Name  string
	Title string
}

func NewRequestContext(update *client.Update, tgClient *client.TgClient, commands []BotCommand) *RequestContext {
	request := &RequestContext{
		TgClient: tgClient,
		Text:     update.Message.Text,
		ChatId:   update.Message.Chat.Id,
		UserInfo: UserInfo{
			LastName:  update.Message.From.LastName,
			UserName:  update.Message.From.Username,
			FirstName: update.Message.From.FirstName,
		},
		Update:   update,
		Commands: commands,
	}

	request.Command = extractCommandName(update)

	if update.Message.CallbackQuery != nil {
		request.CallbackData = update.Message.CallbackQuery.Data
		if request.Command == "" {
			cmdName, btnId := extractInlineButtonData(request)
			request.Command = cmdName
			request.InlineButtonId = btnId
		}
	}

	if update.Message.Document != nil {
		request.FileId = update.Message.Document.FileId
	}

	return request
}

func extractCommandName(update *client.Update) string {
	for _, msgEntity := range update.Message.Entities {
		if msgEntity.Type != "bot_command" {
			continue
		}
		text := update.Message.Text
		if len(text) < msgEntity.Offset+msgEntity.Length {
			return ""
		}
		return text[msgEntity.Offset : msgEntity.Offset+msgEntity.Length]
	}
	return ""
}

func extractInlineButtonData(request *RequestContext) (string, string) {
	query, err := url.ParseQuery(request.CallbackData)
	if err != nil {
		return "", ""
	}
	return query.Get("c"), query.Get("id")
}

func (r *RequestContext) Reply(msg string) (*client.Message, error) {
	return r.ReplyWithButtons(msg, nil, "")
}

func (r *RequestContext) ReplyWithButtons(msg string, inlineButtons []InlineButton, commandName string) (*client.Message, error) {
	params := createMsgParams(r, msg, inlineButtons, commandName)
	return r.TgClient.SendMessage(params)
}

func createMsgParams(r *RequestContext, msg string, inlineButtons []InlineButton, commandName string) *client.SendMessageParams {
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
				CallbackData: "c=" + commandName + "n=" + button.Name,
			}
		}
		params.ReplyMarkup = markup
	}
	return params
}
