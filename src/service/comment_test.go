package service

import (
	"context"
	"fmt"
	"testing"
	"tiktok/src/dal"
	"tiktok/src/handlers"
)

func TestOperateComment(t *testing.T) {
	dal.Init()
	type argument struct {
		ctx context.Context
		req handlers.CommentUpdateParam
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx: context.Background(),
				req: handlers.CommentUpdateParam{
					UserId:     1,
					VideoId:    24,
					ActionType: 1,
					Text:       "嗨嗨嗨我来了",
				},
			}, wantErr: false},
		{
			name: "2", args: argument{
				ctx: context.Background(),
				req: handlers.CommentUpdateParam{
					UserId:     1,
					VideoId:    10,
					ActionType: 2,
					CommentId:  4,
				},
			}, wantErr: false},
		{
			name: "3", args: argument{
				ctx: context.Background(),
				req: handlers.CommentUpdateParam{
					UserId:     999999,
					VideoId:    24,
					ActionType: 1,
					Text:       "嗨嗨嗨我来了",
				},
			}, wantErr: true},
		{
			name: "4", args: argument{
				ctx: context.Background(),
				req: handlers.CommentUpdateParam{
					UserId:     999999,
					VideoId:    10,
					ActionType: 2,
					CommentId:  4,
				},
			}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CommentServiceInstance().OperateComment(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("OperateComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueryCommentByVideoId(t *testing.T) {
	dal.Init()
	type argument struct {
		ctx context.Context
		req handlers.CommentQueryParam
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx: context.Background(),
				req: handlers.CommentQueryParam{
					VideoId: 10,
				},
			}, wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := CommentServiceInstance().QueryCommentByVideoId(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentServiceInstance() error = %v, wantErr = %v", err, tt.wantErr)
			}
			for _, tmp := range res {
				fmt.Println(tmp)
			}

		})
	}
}
