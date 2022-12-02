package client

import (
	"errors"
	"path"
)

// GetMe method for testing your bot's authentication token
func (t *TgClient) GetMe() (user *User, err error) {
	return makeRequest[any, User](t, "getMe", nil)
}

// GetUpdates receives incoming updates using long polling
func (t *TgClient) GetUpdates(params *GetUpdateParams) ([]Update, error) {
	response, err := makeRequest[GetUpdateParams, []Update](t, "getUpdates", params)
	return *response, err
}

// SendMessage sends text messages
func (t *TgClient) SendMessage(params *SendMessageParams) (message *Message, err error) {
	return makeRequest[SendMessageParams, Message](t, "sendMessage", params)
}

// GetFile returns basic information about a file and prepare it for downloading
func (t *TgClient) GetFile(fileId string) (message *File, err error) {
	return makeRequest[getFileParams, File](t, "getFile", &getFileParams{fileId})
}

// SetMyCommands changes the list of the bot's commands
func (t *TgClient) SetMyCommands(params *SetMyCommandsParams) (*bool, error) {
	return makeRequest[SetMyCommandsParams, bool](t, "setMyCommands", params)
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
