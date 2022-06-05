package handlers

type FeedResponse struct {
	Response
	VideoList []FeedInfo `json:"video_list"`
	NextTime  int64      `json:"next_time"`
}

type FeedInfo struct {
	ID           int64    `json:"id"`
	AuthorId     UserInfo `json:"author_id"`
	PlayUrl      string   `json:"play_url"`
	CoverUrl     string   `json:"cover_url"`
	CommentCount int64    `json:"comment_count"`
	Title        string   `json:"title"`
}
