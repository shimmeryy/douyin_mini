package db

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateFavorInfo(t *testing.T) {
	Init()
	type argument struct {
		ctx   context.Context
		favor Favor
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx: context.Background(),
				favor: Favor{
					UserId:  1,
					VideoId: 10,
				},
			}, wantErr: false},
		{
			name: "2", args: argument{
				ctx: context.Background(),
				favor: Favor{
					UserId:  2,
					VideoId: 15,
				},
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateFavorInfo(tt.args.ctx, tt.args.favor); (err != nil) != tt.wantErr {
				t.Errorf("CreateFavorInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueryFavorByUserId(t *testing.T) {
	Init()
	type argument struct {
		ctx    context.Context
		userId int64
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx:    context.Background(),
				userId: 1,
			}, wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := QueryFavorByUserId(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryFavorByUserId() error = %v, wantErr = %v", err, tt.wantErr)
			}
			for _, tmp := range res {
				fmt.Println(tmp)
			}

		})
	}
}

func TestDeleteFavorInfo(t *testing.T) {
	Init()
	type argument struct {
		ctx     context.Context
		userId  int64
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
				userId:  1,
				videoId: 15,
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteFavorInfo(tt.args.ctx, tt.args.userId, tt.args.videoId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFavorInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
