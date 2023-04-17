package durov

import (
	"fmt"
	"github.com/ayrat404/durov/client"
	"net/url"
	"strings"
)

// RequestContext represents an incoming request from a user
type RequestContext struct {
	Command      string
	CommandData  map[string]string
	Text         string
	CallbackData string
	FileId       string
	ChatId       int
	Update       *client.Update
	TgClient     *client.TgClient
	UserInfo     UserInfo
}

// UserInfo contains user info
type UserInfo struct {
	UserName  string
	FirstName string
	LastName  string
}

type InlineButton struct {
	Title string
	Args  map[string]string
}

func NewRequestContext(update *client.Update, tgClient *client.TgClient, commands []BotCommand) *RequestContext {
	request := &RequestContext{
		TgClient: tgClient,
		Update:   update,
	}

	if update.Message != nil {
		request.Text = update.Message.Text
		request.ChatId = update.Message.Chat.Id
		request.UserInfo = UserInfo{
			LastName:  update.Message.From.LastName,
			UserName:  update.Message.From.Username,
			FirstName: update.Message.From.FirstName,
		}
		request.Command = extractCommandName(update)
		if update.Message.Document != nil {
			request.FileId = update.Message.Document.FileId
		}
	}

	if update.CallbackQuery != nil {
		request.CallbackData = update.CallbackQuery.Data
		if request.Command == "" {
			cmdName, cmdData := extractInlineButtonData(request)
			request.Command = cmdName
			request.CommandData = cmdData
		}
		request.UserInfo = UserInfo{
			LastName:  update.CallbackQuery.From.LastName,
			UserName:  update.CallbackQuery.From.Username,
			FirstName: update.CallbackQuery.From.FirstName,
		}
		if update.CallbackQuery.Message != nil {
			request.ChatId = update.CallbackQuery.Message.Chat.Id
		}
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

func extractInlineButtonData(request *RequestContext) (string, map[string]string) {
	data := make(map[string]string)
	query, err := url.ParseQuery(request.CallbackData)
	if err != nil {
		return "", data
	}
	cmd := ""
	for k, v := range query {
		if k == "c" {
			cmd = v[0]
			continue
		}
		data[k] = v[0]
	}
	return cmd, data
}

func (r *RequestContext) Reply(msg string) (*client.Message, error) {
	return r.ReplyWithButtons(msg, nil, "")
}

func (r *RequestContext) ReplyWithButtons(msg string, inlineButtons [][]InlineButton, commandName string) (*client.Message, error) {
	params := createMsgParams(r, msg, inlineButtons, commandName)
	return r.TgClient.SendMessage(params)
}

func createMsgParams(r *RequestContext, msg string, inlineButtons [][]InlineButton, commandName string) *client.SendMessageParams {
	params := &client.SendMessageParams{
		Text:   msg,
		ChatId: r.ChatId,
	}
	if len(inlineButtons) > 0 {
		markup := client.InlineKeyboardMarkup{
			InlineKeyboard: make([][]client.InlineKeyboardButton, len(inlineButtons), len(inlineButtons)),
		}
		for i, buttonRow := range inlineButtons {
			markup.InlineKeyboard[i] = make([]client.InlineKeyboardButton, len(buttonRow), len(buttonRow))
			for j, button := range buttonRow {
				markup.InlineKeyboard[i][j] = client.InlineKeyboardButton{
					Text:         button.Title,
					CallbackData: "c=" + commandName + "&" + queryStringFromMap(button.Args),
				}
			}
		}
		params.ReplyMarkup = markup
	}
	return params
}

func queryStringFromMap(m map[string]string) string {
	sb := strings.Builder{}
	for k, v := range m {
		if sb.Len() > 0 {
			sb.WriteString("&")
		}
		fmt.Fprintf(&sb, "%s=%s", k, v)
	}
	return sb.String()
}
