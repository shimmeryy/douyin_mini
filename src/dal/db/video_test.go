package db

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateVideo(t *testing.T) {
	Init()
	type args struct {
		ctx   context.Context
		video Video
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "1", args: args{
			ctx: context.Background(),
			video: Video{
				AuthorId:     1,
				PlayUrl:      "this is a test",
				CoverUrl:     "this is a test",
				CommentCount: 0,
				Title:        "this is a test",
			},
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateVideo(tt.args.ctx, tt.args.video); (err != nil) != tt.wantErr {
				t.Errorf("CreateVideo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueryVideoByAuthor(t *testing.T) {
	Init()
	type args struct {
		ctx      context.Context
		authorId int64
	}
	tests := []struct {
		name    string
		args    args
		want    []*Video
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				ctx:      context.Background(),
				authorId: 1,
			}, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryVideoByAuthor(tt.args.ctx, tt.args.authorId)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryVideoByAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, res := range got {
				fmt.Println(res)
			}
		})
	}
}

func TestQueryVideoById(t *testing.T) {
	Init()
	type args struct {
		ctx     context.Context
		videoId int64
	}
	tests := []struct {
		name    string
		args    args
		want    *Video
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				ctx:     context.Background(),
				videoId: 1,
			}, wantErr: false,
		},
		{
			name: "2",
			args: args{
				ctx:     context.Background(),
				videoId: 2,
			}, wantErr: false,
		},
		{
			name: "3",
			args: args{
				ctx:     context.Background(),
				videoId: 3,
			}, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryVideoById(tt.args.ctx, tt.args.videoId)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryVideoById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}
