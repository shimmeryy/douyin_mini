package handlers

import "tiktok/src/dal/db"

type FavorOperateParam struct {
	UserId     int64 `json:"user_id"`
	VideoId    int64 `json:"video_id"`
	ActionType int32 `json:"action_type"`
}

type FavorQueryParam struct {
	UserId int64 `json:"user_id"`
}

type FavorQueryResponse struct {
	Response
	VideoList []*db.Video
}

type FavorCountParam struct {
	VideoId int64 `json:"video_id"`
}

type FavorCheckParam struct {
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}
