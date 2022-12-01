package client

import (
	"errors"
	"path"
)

func (t *TgClient) GetMe() (user *User, err error) {
	return makeRequest[any, User](t, "getMe", nil)
}

func (t *TgClient) GetUpdates(params *GetUpdateParams) ([]Update, error) {
	response, err := makeRequest[GetUpdateParams, []Update](t, "getUpdates", params)
	return *response, err
}

func (t *TgClient) SendMessage(params *SendMessageParams) (message *Message, err error) {
	return makeRequest[SendMessageParams, Message](t, "sendMessage", params)
}

func (t *TgClient) GetFile(fileId string) (message *File, err error) {
	return makeRequest[GetFileParams, File](t, "getFile", &GetFileParams{fileId})
}

func (t *TgClient) SetMyCommands(params *SetMyCommandsParams) (*bool, error) {
	return makeRequest[SetMyCommandsParams, bool](t, "setMyCommands", params)
}

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
