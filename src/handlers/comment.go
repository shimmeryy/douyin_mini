package handlers

type CommentOperateParam struct {
	UserId     int64  `json:"user_id"`
	VideoId    int64  `json:"video_id"`
	ActionType int32  `json:"action_type"`
	Text       string `json:"text"`
	CommentId  int64  `json:"comment_id"`
}

type CommentQueryParam struct {
	VideoId int64 `json:"video_id"`
}
