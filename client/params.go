package client

// GetUpdateParams parameters for TgClient.GetUpdates method
type GetUpdateParams struct {
	Offset         int      `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// SendMessageParams parameters for TgClient.SendMessage method
type SendMessageParams struct {
	ChatId      int    `json:"chat_id"`
	Text        string `json:"text"`
	ReplyMarkup any    `json:"reply_markup,omitempty"` // only InlineKeyboardMarkup is supported for now
}

type InlineKeyboardMarkup struct {
	InlineKeyboard []InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	Url          string `json:"url,omitempty"`
	CallbackData string `json:"callback_data,omitempty"`
}

// getFileParams parameters for TgClient.GetFile method
type getFileParams struct {
	FileId string `json:"file_id"`
}

// SetMyCommandsParams parameters for TgClient.SetMyCommands method
type SetMyCommandsParams struct {
	Commands     []BotCommand `json:"commands"`
	LanguageCode string       `json:"language_code,omitempty"`
}
