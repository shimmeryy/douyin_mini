package handlers

type FavorUpdateParam struct {
	UserId     int64 `json:"user_id"`
	VideoId    int64 `json:"video_id"`
	ActionType int32 `json:"action_type"`
}

type FavorQueryParam struct {
	UserId int64 `json:"user_id"`
}
