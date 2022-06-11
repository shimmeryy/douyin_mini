package db

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateCommentInfo(t *testing.T) {
	Init()
	type argument struct {
		ctx     context.Context
		comment Comment
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx: context.Background(),
				comment: Comment{
					UserId:  1,
					VideoId: 10,
					Text:    "test",
				},
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := CreateCommentInfo(tt.args.ctx, tt.args.comment); (err != nil) != tt.wantErr {
				t.Errorf("CreateCommentInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueryCommentByVideoId(t *testing.T) {
	Init()
	type argument struct {
		ctx     context.Context
		videoId int64
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx:     context.Background(),
				videoId: 10,
			}, wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := QueryCommentByVideoId(tt.args.ctx, tt.args.videoId)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryCommentByVideoId() error = %v, wantErr = %v", err, tt.wantErr)
			}
			for _, tmp := range res {
				fmt.Println(tmp)
			}

		})
	}
}

func TestQueryCommentById(t *testing.T) {
	Init()
	type argument struct {
		ctx       context.Context
		commentId int64
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx:       context.Background(),
				commentId: 6,
			}, wantErr: false,
		},
		{
			name: "2", args: argument{
				ctx:       context.Background(),
				commentId: 12,
			}, wantErr: false,
		},
		{
			name: "3", args: argument{
				ctx:       context.Background(),
				commentId: 9999,
			}, wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := QueryCommentById(tt.args.ctx, tt.args.commentId)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryCommentById() error = %v, wantErr = %v", err, tt.wantErr)
			}
			fmt.Println(res)
		})
	}
}

func TestCheckCommentById(t *testing.T) {
	Init()
	type argument struct {
		ctx       context.Context
		commentId int64
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx:       context.Background(),
				commentId: 6,
			}, wantErr: false,
		},
		{
			name: "2", args: argument{
				ctx:       context.Background(),
				commentId: 12,
			}, wantErr: false,
		},
		{
			name: "3", args: argument{
				ctx:       context.Background(),
				commentId: 9999,
			}, wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := CheckCommentById(tt.args.ctx, tt.args.commentId)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckCommentById() error = %v, wantErr = %v", err, tt.wantErr)
			}
			fmt.Println(res)
		})
	}
}

func TestDeleteCommentInfo(t *testing.T) {
	Init()
	type argument struct {
		ctx       context.Context
		commentId int64
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx:       context.Background(),
				commentId: 1,
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteCommentInfo(tt.args.ctx, tt.args.commentId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCommentInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCountCommentByVideoId(t *testing.T) {
	Init()
	type argument struct {
		ctx     context.Context
		videoId int64
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx:     context.Background(),
				videoId: 10,
			}, wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cnt, err := CountCommentByVideoId(tt.args.ctx, tt.args.videoId)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountCommentByVideoId() error = %v, wantErr = %v", err, tt.wantErr)
			}
			fmt.Println(cnt)
		})
	}
}
