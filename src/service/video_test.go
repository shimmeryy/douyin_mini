package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"tiktok/src/dal"
	"tiktok/src/dal/db"
	"tiktok/src/handlers"
)

func TestCheckVideoById(t *testing.T) {
	dal.Init()
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{
			ctx: context.Background(),
			id:  1,
		}, want: true,
		},
		{name: "1", args: args{
			ctx: context.Background(),
			id:  2,
		}, want: true,
		},
		{name: "1", args: args{
			ctx: context.Background(),
			id:  10086,
		}, want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &VideoServiceImpl{}
			if got := this.CheckVideoById(tt.args.ctx, tt.args.id); got != tt.want {
				t.Errorf("CheckVideoById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetVideoById(t *testing.T) {
	dal.Init()
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name          string
		args          args
		wantVideoInfo *handlers.VideoInfo
		wantErr       bool
	}{
		{
			name: "1",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
		},
		{
			name: "2",
			args: args{
				ctx: context.Background(),
				id:  2,
			},
		},
		{
			name: "3",
			args: args{
				ctx: context.Background(),
				id:  10086,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &VideoServiceImpl{}
			gotVideoInfo, err := this.GetVideoById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVideoById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVideoInfo, tt.wantVideoInfo) {
				t.Errorf("GetVideoById() gotVideoInfo = %v, want %v", gotVideoInfo, tt.wantVideoInfo)
			}
		})
	}
}

func TestGetVideosByAuthor(t *testing.T) {
	dal.Init()
	type args struct {
		ctx    context.Context
		userId int64
	}
	tests := []struct {
		name     string
		args     args
		wantList []*db.Video
		wantErr  bool
	}{
		{
			name: "1",
			args: args{
				ctx:    context.Background(),
				userId: 1,
			},
		},
		{
			name: "2",
			args: args{
				ctx:    context.Background(),
				userId: 10086,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &VideoServiceImpl{}
			gotList, err := this.GetVideosByAuthor(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVideosByAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i := range gotList {
				fmt.Println(gotList[i])
			}
		})
	}
}
