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

func TestRegisterUser(t *testing.T) {
	dal.Init()
	type argument struct {
		ctx      context.Context
		username string
		password string
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx:      context.Background(),
				username: "test01",
				password: "",
			}, wantErr: true},
		{
			name: "2", args: argument{
				ctx:      context.Background(),
				username: "test02",
				password: "000000000000000000000000000000000000000000000000000",
			}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := UserServiceInstance().RegisterUser(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserFollowers(t *testing.T) {
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
				ID:  10,
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := UserServiceInstance().GetUserFollowers(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserFollowers() error = %v, wantErr %v", err, tt.wantErr)
			}
			for _, tmp := range res {
				fmt.Println(tmp)
			}
		})
	}
}

func TestGetUserFans(t *testing.T) {
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
				ID:  2,
			}, wantErr: false},
		{
			name: "2", args: argument{
				ctx: context.Background(),
				ID:  3,
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := UserServiceInstance().GetUserFans(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserFans() error = %v, wantErr %v", err, tt.wantErr)
			}
			for _, tmp := range res {
				fmt.Println(tmp)
			}
		})
	}
}

func TestFollowUser(t *testing.T) {
	dal.Init()
	type argument struct {
		ctx          context.Context
		userId       int64
		followUserId int64
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx:          context.Background(),
				userId:       1,
				followUserId: 2,
			}, wantErr: false},
		{
			name: "2", args: argument{
				ctx:          context.Background(),
				userId:       1,
				followUserId: 3,
			}, wantErr: false},
		{
			name: "3", args: argument{
				ctx:          context.Background(),
				userId:       2,
				followUserId: 3,
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UserServiceInstance().FollowUser(tt.args.ctx, tt.args.userId, tt.args.followUserId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FollowUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCancelFollowUser(t *testing.T) {
	dal.Init()
	type argument struct {
		ctx          context.Context
		userId       int64
		followUserId int64
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx:          context.Background(),
				userId:       1,
				followUserId: 2,
			}, wantErr: false},
		{
			name: "2", args: argument{
				ctx:          context.Background(),
				userId:       1,
				followUserId: 3,
			}, wantErr: false},
		{
			name: "3", args: argument{
				ctx:          context.Background(),
				userId:       2,
				followUserId: 3,
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UserServiceInstance().CancelFollowUser(tt.args.ctx, tt.args.userId, tt.args.followUserId)
			if (err != nil) != tt.wantErr {
				t.Errorf("CancelFollowUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
