package client

import (
	"encoding/json"
	"fmt"
)

type TgResponse struct {
	Ok          bool            `json:"ok"`
	Description string          `json:"description"`
	ErrorCode   int             `json:"error_code"`
	Result      json.RawMessage `json:"result"`
}

type TgError struct {
	Description string
	ErrorCode   int
	//HttpStatusCode string
	//RawResponse    string
}

func (t *TgError) Error() string {
	return t.Description
}

func (t *TgError) String() string {
	return fmt.Sprintf("error code: %v, description: %v", t.ErrorCode, t.Description)
}

type User struct {
	Id                      int    `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Username                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	IsPremium               bool   `json:"is_premium"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

type Update struct {
	UpdateId int      `json:"update_id"`
	Message  *Message `json:"message,omitempty"`
}

type Message struct {
	MessageId int             `json:"message_id"`
	From      User            `json:"from,omitempty"`
	Chat      Chat            `json:"chat"`
	Date      int             `json:"date"`
	Text      string          `json:"text,omitempty"`
	Document  *Document       `json:"document,omitempty"`
	Caption   string          `json:"caption,omitempty"`
	Photo     []PhotoSize     `json:"photo,omitempty"`
	Entities  []MessageEntity `json:"entities,omitempty"`
}

type Chat struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type Document struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int    `json:"file_size,omitempty"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}

type File struct {
	FileId       string `json:"file_Id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int    `json:"file_size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}

type DownloadedFile struct {
	Content []byte
	Name    string
}

type MessageEntity struct {
	Offset   int    `json:"offset"`
	Length   int    `json:"length"`
	Type     string `json:"type"`
	Url      string `json:"url,omitempty"`
	User     *User  `json:"user,omitempty"`
	Language string `json:"language,omitempty"`
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}
