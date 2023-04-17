package client

import (
	"context"
	"errors"
	"path"
)

// GetMe method for testing your bot's authentication token
func (t *TgClient) GetMe(ctx context.Context) (user *User, err error) {
	return makeRequest[any, User](t, ctx, "getMe", nil)
}

// GetUpdates receives incoming updates using long polling
func (t *TgClient) GetUpdates(ctx context.Context, params *GetUpdateParams) ([]Update, error) {
	response, err := makeRequest[GetUpdateParams, []Update](t, ctx, "getUpdates", params)
	if err != nil {
		return nil, err
	}
	return *response, nil
}

// SendMessage sends text messages
func (t *TgClient) SendMessage(params *SendMessageParams) (message *Message, err error) {
	return makeRequest[SendMessageParams, Message](t, context.Background(), "sendMessage", params)
}

// GetFile returns basic information about a file and prepare it for downloading
func (t *TgClient) GetFile(fileId string) (message *File, err error) {
	return makeRequest[getFileParams, File](t, context.Background(), "getFile", &getFileParams{fileId})
}

// SetMyCommands changes the list of the bot's commands
func (t *TgClient) SetMyCommands(params *SetMyCommandsParams) (*bool, error) {
	return makeRequest[SetMyCommandsParams, bool](t, context.Background(), "setMyCommands", params)
}

// DownloadFile downloads a file
func (t *TgClient) DownloadFile(fileId string) (*DownloadedFile, error) {
	file, err := t.GetFile(fileId)
	if err != nil {
		return nil, err
	}
	if file.FilePath == "" {
		return nil, errors.New("empty filePath")
	}

	content, err := downloadFile(t, file.FilePath)
	if err != nil {
		return nil, err
	}

	return &DownloadedFile{content, path.Base(file.FilePath)}, nil
}

type AnswerCallbackQueryParams struct {
	CallbackQueryId string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	Url             string `json:"url,omitempty"`
	CacheTime       string `json:"cache_time,omitempty"`
}

func (t *TgClient) AnswerCallbackQuery(params *AnswerCallbackQueryParams) (bool, error) {
	result, err := makeRequest[AnswerCallbackQueryParams, bool](t, context.Background(), "answerCallbackQuery", params)
	return *result, err
}
