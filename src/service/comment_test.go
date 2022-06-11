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
		req handlers.CommentOperateParam
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx: context.Background(),
				req: handlers.CommentOperateParam{
					UserId:     1,
					VideoId:    15,
					ActionType: 1,
					Text:       "test",
				},
			}, wantErr: false},
		{
			name: "2", args: argument{
				ctx: context.Background(),
				req: handlers.CommentOperateParam{
					UserId:     1,
					VideoId:    10,
					ActionType: 2,
					CommentId:  7,
				},
			}, wantErr: false},
		{
			name: "3", args: argument{
				ctx: context.Background(),
				req: handlers.CommentOperateParam{
					UserId:     999999,
					VideoId:    15,
					ActionType: 1,
					Text:       "test",
				},
			}, wantErr: true},
		{
			name: "4", args: argument{
				ctx: context.Background(),
				req: handlers.CommentOperateParam{
					UserId:     999999,
					VideoId:    10,
					ActionType: 2,
					CommentId:  4,
				},
			}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := CommentServiceInstance().OperateComment(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("OperateComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueryCommentById(t *testing.T) {
	dal.Init()
	type argument struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx: context.Background(),
				id:  12,
			}, wantErr: false,
		},
		{
			name: "2", args: argument{
				ctx: context.Background(),
				id:  9999999,
			}, wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := CommentServiceInstance().QueryCommentById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryCommentById() error = %v, wantErr = %v", err, tt.wantErr)
			}
			fmt.Println(res)
		})
	}
}

func TestQueryCommentByVideoId(t *testing.T) {
	dal.Init()
	type argument struct {
		ctx context.Context
		req handlers.CommentQueryByVideoIdParam
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx: context.Background(),
				req: handlers.CommentQueryByVideoIdParam{
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

func TestCountCommentByVideoId(t *testing.T) {
	dal.Init()
	type argument struct {
		ctx context.Context
		req handlers.CommentQueryByVideoIdParam
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx: context.Background(),
				req: handlers.CommentQueryByVideoIdParam{
					VideoId: 10,
				},
			}, wantErr: false,
		},
		{
			name: "2", args: argument{
				ctx: context.Background(),
				req: handlers.CommentQueryByVideoIdParam{
					VideoId: 9999,
				},
			}, wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := CommentServiceInstance().CountCommentByVideoId(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountCommentByVideoId() error = %v, wantErr = %v", err, tt.wantErr)
			}
			fmt.Println(res)
		})
	}
}
