package handlers

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
	UserName      string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type UserFollowersResponse struct {
	Response
	Followers []UserInfo `json:"user_list"`
}

type UserFansResponse struct {
	Response
	Fans []UserInfo `json:"user_list"`
}

type UserRegisterResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}
