package handlers

//登录
type UserLoginParam struct {
	UserName string `json:"username" form:"username" `
	PassWord string `json:"password" form:"password"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

type UserInfoResponse struct {
	Response
	User UserInfo `json:"user"`
}

type UserInfo struct {
	ID            int64  `json:"id"`
	UserName      string `json:"user_name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
