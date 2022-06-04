package db

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateComment(t *testing.T) {
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
					Text:    "haha",
				},
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateCommentInfo(tt.args.ctx, tt.args.comment); (err != nil) != tt.wantErr {
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
