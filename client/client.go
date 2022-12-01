package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const tgHost = "https://api.telegram.org/"

type TgClient struct {
	token  string
	client *http.Client
}

func NewClient(token string) *TgClient {
	return &TgClient{token, &http.Client{}}
}

func makeRequest[Req any, Resp any](t *TgClient, method string, params *Req) (response *Resp, err error) {
	var bodyReader io.Reader
	if params != nil {
		bodyJson, err := json.Marshal(params)
		if err != nil {
			return response, err
		}
		bodyReader = bytes.NewReader(bodyJson)
	}

	rawResp, err := t.client.Post(formatUrl(t.token, method), "application/json", bodyReader)
	if err != nil {
		return nil, err
	}

	defer rawResp.Body.Close()

	return processResponse[Resp](rawResp.Body)
}

func processResponse[Resp any](resp io.Reader) (*Resp, error) {
	raw, err := io.ReadAll(resp)
	if err != nil {
		return nil, err
	}

	var apiResp TgResponse
	err = json.Unmarshal(raw, &apiResp)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, &TgError{apiResp.Description, apiResp.ErrorCode}
	}

	var response Resp
	err = json.Unmarshal(apiResp.Result, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func downloadFile(t *TgClient, filePath string) ([]byte, error) {
	fileUrl := formatFileUrl(t.token, filePath)
	rawResp, err := t.client.Get(fileUrl)
	if err != nil {
		return nil, err
	}

	defer rawResp.Body.Close()

	if rawResp.StatusCode != 200 {
		return nil, errors.New("failed to download file: http status code " + strconv.Itoa(rawResp.StatusCode))
	}

	return io.ReadAll(rawResp.Body)
}

func formatUrl(token string, path string) string {
	return fmt.Sprintf("%sbot%s/%s", tgHost, token, path)
}

func formatFileUrl(token string, filePath string) string {
	return fmt.Sprintf("%sfile/bot%s/%s", tgHost, token, filePath)
}
