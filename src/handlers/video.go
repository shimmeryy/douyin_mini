package handlers

import (
	"mime/multipart"
)

type UploadResponse struct {
	Response
}

type PublishInfo struct {
	ID            int64    `json:"id"`
	Author        UserInfo `json:"author"`
	PlayUrl       string   `json:"play_url"`
	CoverUrl      string   `json:"cover_url"`
	FavoriteCount int64    `json:"favorite_count"`
	IsFavorite    bool     `json:"is_favorite"`
	CommentCount  int64    `json:"comment_count"`
	Title         string   `json:"title"`
}
type VideoInfo struct {
	ID            int64    `json:"id"`
	Author        UserInfo `json:"author"`
	PlayUrl       string   `json:"play_url"`
	CoverUrl      string   `json:"cover_url"`
	FavoriteCount int64    `json:"favorite_count"`
	IsFavorite    bool     `json:"is_favorite"`
	CommentCount  int64    `json:"comment_count"`
	Title         string   `json:"title"`
}
type PublishParams struct {
	Data   *multipart.FileHeader `json:"data"`
	UserId int64                 `json:"user_id"`
	Title  string                `json:"title"`
}

type PublishListResponse struct {
	Response
	VideoList []VideoInfo `json:"video_list"`
}
