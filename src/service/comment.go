package service

import (
	"context"
	"sync"
	"tiktok/src/dal/db"
	"tiktok/src/errno"
	"tiktok/src/handlers"
)

type CommentService interface {
	CreateComment()
	QueryCommentByVideoId()
	QueryCommentById()
	CountCommentByVideoId()
}

var (
	commentService     *CommentServiceImpl
	commentServiceOnce sync.Once
)

func CommentServiceInstance() *CommentServiceImpl {
	commentServiceOnce.Do(
		func() {
			commentService = &CommentServiceImpl{}
		})
	return commentService
}

type CommentServiceImpl struct {
}

func (this *CommentServiceImpl) OperateComment(ctx context.Context, req handlers.CommentOperateParam) (*handlers.CommentInfo, error) {
	// check UserId is valid or not
	if flag := UserServiceInstance().CheckUserById(ctx, req.UserId); flag != true {
		return nil, errno.ServiceErr.WithMessage("invalid userId")
	}

	// check VideoId is valid or not
	if flag := VideoServiceInstance().CheckVideoById(ctx, req.VideoId); flag != true {
		return nil, errno.ServiceErr.WithMessage("invalid videoId")
	}

	if req.ActionType != 1 && req.ActionType != 2 {
		return nil, errno.ServiceErr.WithMessage("invalid actionType")
	}

	// ActionType == 2 delete comment
	if req.ActionType == 2 {
		flag, _ := db.CheckCommentById(ctx, req.CommentId) // 校验commentId是否合法
		if !flag {
			return nil, errno.ServiceErr.WithMessage("invalid commentId")
		}
		if err := db.DeleteCommentInfo(ctx, req.CommentId); err != nil {
			return nil, err
		}
		return nil, nil
	}
	// ActionType == 1 create comment
	comment := db.Comment{
		UserId:  req.UserId,
		VideoId: req.VideoId,
		Text:    req.Text,
	}
	commentId, err := db.CreateCommentInfo(ctx, comment)
	if err != nil {
		return nil, err
	}
	res, err := this.QueryCommentById(ctx, commentId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (this *CommentServiceImpl) QueryCommentById(ctx context.Context, id int64) (*handlers.CommentInfo, error) {
	comment, err := db.QueryCommentById(ctx, id)
	if err != nil {
		return nil, err
	}

	userId := comment.UserId
	userInfo, err := UserServiceInstance().GetUserInfo(ctx, userId)
	if err != nil {
		return nil, err
	}

	commentInfo := handlers.CommentInfo{
		CommentId:  int64(comment.ID),
		UserInfo:   *userInfo,
		Content:    comment.Text,
		CreateDate: comment.CreatedAt.Format("01-02"),
	}

	return &commentInfo, nil
}

func (this *CommentServiceImpl) QueryCommentByVideoId(ctx context.Context, req handlers.CommentQueryByVideoIdParam) ([]*handlers.CommentInfo, error) {
	if flag := VideoServiceInstance().CheckVideoById(ctx, req.VideoId); flag != true {
		return nil, errno.ServiceErr.WithMessage("invalid videoId")
	}

	var res []*handlers.CommentInfo
	commentList, err := db.QueryCommentByVideoId(ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	for _, comment := range commentList {
		userId := comment.UserId
		userInfo, err := UserServiceInstance().GetUserInfo(ctx, userId)
		if err != nil {
			return nil, err
		}
		commentInfo := handlers.CommentInfo{
			CommentId:  int64(comment.ID),
			UserInfo:   *userInfo,
			Content:    comment.Text,
			CreateDate: comment.CreatedAt.Format("01-02"),
		}
		res = append(res, &commentInfo)
	}

	return res, nil
}

func (this *CommentServiceImpl) CountCommentByVideoId(ctx context.Context, req handlers.CommentQueryByVideoIdParam) (int64, error) {
	cnt, err := db.CountCommentByVideoId(ctx, req.VideoId)
	if err != nil {
		return cnt, err
	}
	return cnt, nil
}
