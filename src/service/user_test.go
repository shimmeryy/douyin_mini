package service

import (
	"context"
	"fmt"
	"testing"
	"tiktok/src/dal"
	"tiktok/src/handlers"
)

func TestCheckUser(t *testing.T) {
	dal.Init()
	type argument struct {
		ctx context.Context
		req handlers.UserLoginParam
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx: context.Background(),
				req: handlers.UserLoginParam{
					UserName: "changlu",
					PassWord: "123456",
				},
			}, wantErr: false},
		{
			name: "2", args: argument{
				ctx: context.Background(),
				req: handlers.UserLoginParam{
					UserName: "changlu",
					PassWord: "12345",
				},
			}, wantErr: false},
		{
			name: "3", args: argument{
				ctx: context.Background(),
				req: handlers.UserLoginParam{
					UserName: "changlu1",
					PassWord: "123456",
				},
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := UserServiceInstance().CheckUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println(res)
		})
	}
}

func TestGetUserInfo(t *testing.T) {
	dal.Init()
	type argument struct {
		ctx context.Context
		ID  int64
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx: context.Background(),
				ID:  1,
			}, wantErr: false},
		{
			name: "2", args: argument{
				ctx: context.Background(),
				ID:  2,
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := UserServiceInstance().GetUserInfo(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println(res)
		})
	}
}

func TestCheckUserById(t *testing.T) {
	dal.Init()
	type argument struct {
		ctx context.Context
		ID  int64
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx: context.Background(),
				ID:  1,
			}, wantErr: false},
		{
			name: "2", args: argument{
				ctx: context.Background(),
				ID:  2,
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := UserServiceInstance().CheckUserById(tt.args.ctx, tt.args.ID)
			fmt.Println(res)
		})
	}
}
