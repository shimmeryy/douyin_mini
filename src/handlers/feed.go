package handlers

type FeedResponse struct {
	Response
	NextTime  int64      `json:"next_time"`
	VideoList []FeedInfo `json:"video_list"`
}

type FeedInfo struct {
	ID           int64    `json:"id"`
	Author       UserInfo `json:"author"`
	PlayUrl      string   `json:"play_url"`
	CoverUrl     string   `json:"cover_url"`
	IsFavorite   bool     `json:"is_favorite"`
	CommentCount int64    `json:"comment_count"`
	Title        string   `json:"title"`
}
