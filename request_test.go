package durov

import (
	"github.com/ayrat404/durov/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRequestContext(t *testing.T) {
	update := buildUpdate()

	req := NewRequestContext(update, &client.TgClient{}, []BotCommand{})

	assert.NotNil(t, req)
	assert.NotNil(t, req.TgClient)
	assert.NotNil(t, req.Commands)
	assert.NotNil(t, req.Update)

	assert.Equal(t, "text /command1", req.Text)
	assert.Equal(t, "fileId123", req.FileId)
	assert.Equal(t, 123, req.ChatId)

	assert.Equal(t, "someName", req.UserInfo.UserName)
	assert.Equal(t, "Firstname", req.UserInfo.FirstName)
	assert.Equal(t, "Lastname", req.UserInfo.LastName)

	assert.Equal(t, "/command1", req.Command)
	assert.Equal(t, "c=/command2&id=321", req.CallbackData)
	assert.Empty(t, req.InlineButtonId)
}

func TestNewRequestContextCommandFromCallbackQuery(t *testing.T) {
	update := buildUpdate()
	update.Message.Entities = []client.MessageEntity{}

	req := NewRequestContext(update, &client.TgClient{}, []BotCommand{})

	assert.Equal(t, "/command2", req.Command)
	assert.Equal(t, "321", req.InlineButtonId)
}

func buildUpdate() *client.Update {
	return &client.Update{
		UpdateId: 123,
		Message: &client.Message{
			Text: "text /command1",
			CallbackQuery: &client.CallbackQuery{
				Data: "c=/command2&id=321",
			},
			Chat: client.Chat{
				Id: 123,
			},
			Document: &client.Document{
				FileId: "fileId123",
			},
			From: client.User{
				Username:  "someName",
				FirstName: "Firstname",
				LastName:  "Lastname",
			},
			Entities: []client.MessageEntity{
				{
					Length: 9,
					Type:   "bot_command",
					Offset: 5,
				},
			},
		},
	}
}
