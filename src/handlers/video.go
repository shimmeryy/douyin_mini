package handlers

import (
	"mime/multipart"
)

type UploadResponse struct {
	Response
}
type VideoInfo struct {
	ID           int64    `json:"id"`
	Author       UserInfo `json:"author"`
	PlayUrl      string   `json:"play_url"`
	CoverUrl     string   `json:"cover_url"`
	CommentCount int64    `json:"comment_count"`
	Title        string   `json:"title"`
}
type PublishParams struct {
	Data   *multipart.FileHeader `json:"data"`
	UserId int64                 `json:"user_id"`
	Title  string                `json:"title"`
}
