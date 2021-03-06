package handlers

type CommentOperateParam struct {
	UserId     int64  `json:"user_id"`
	VideoId    int64  `json:"video_id"`
	ActionType int32  `json:"action_type"`
	Text       string `json:"text"`
	CommentId  int64  `json:"comment_id"`
}

type CommentQueryByVideoIdParam struct {
	VideoId int64 `json:"video_id"`
}

type CommentInfo struct {
	CommentId  int64 `json:"id"`
	UserInfo   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type CommentCreateResponse struct {
	Response
	CommentInfo `json:"comment"`
}

type CommentQueryResponse struct {
	Response
	CommentList []*CommentInfo `json:"comment_list"`
}
