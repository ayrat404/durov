package client

// GetUpdateParams parameters for getUpdates method
type GetUpdateParams struct {
	Offset         int      `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

type SendMessageParams struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

type GetFileParams struct {
	FileId string `json:"file_id"`
}

type SetMyCommandsParams struct {
	Commands     []BotCommand `json:"commands"`
	LanguageCode string       `json:"language_code,omitempty"`
}
