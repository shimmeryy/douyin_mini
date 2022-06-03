package service

import (
	"context"
	"errors"
	"sync"
	"tiktok/src/dal/db"
	"tiktok/src/handlers"
)

type CommentService interface {
	CreateComment()
	QueryCommentByVideoId()
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

func (this *CommentServiceImpl) OperateComment(ctx context.Context, req handlers.CommentUpdateParam) error {
	// check UserId is valid or not
	if flag := UserServiceInstance().CheckUserById(ctx, req.UserId); flag != true {
		return errors.New("invalid userId")
	}

	// TODO 检验videoId是否合法

	// ActionType == 2 delete comment
	if req.ActionType == 2 {
		if err := db.DeleteCommentInfo(ctx, req.CommentId); err != nil {
			return err
		}
		return nil
	}
	// ActionType == 1 create comment
	comment := db.Comment{
		UserId:  req.UserId,
		VideoId: req.VideoId,
		Text:    req.Text,
	}
	if err := db.CreateCommentInfo(ctx, comment); err != nil {
		return err
	}
	return nil
}

func (this *CommentServiceImpl) QueryCommentByVideoId(ctx context.Context, req handlers.CommentQueryParam) ([]*db.Comment, error) {
	var res []*db.Comment
	res, err := db.QueryCommentByVideoId(ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
