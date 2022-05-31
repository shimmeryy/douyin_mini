package db

import (
	"context"
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
		// TODO: Add test cases.
		{name: "1", args: args{
			ctx: context.Background(),
			video: Video{
				AuthorId:     1,
				PlayUrl:      "hello",
				CoverUrl:     "hello",
				CommentCount: 0,
				Title:        "ok",
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
		{name: "1", args: args{
			ctx:      context.Background(),
			authorId: 1,
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryVideoByAuthor(tt.args.ctx, tt.args.authorId)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryVideoByAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v", got)
		})
	}
}
