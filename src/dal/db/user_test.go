package db

import (
	"context"
	"fmt"
	"testing"
)

func TestQueryUser(t *testing.T)  {
	Init()
	type argument struct {
		ctx     context.Context
		userName string
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
			ctx: context.Background(),
			userName: "changlu",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := QueryUser(tt.args.ctx, tt.args.userName)
			if  (err != nil) != tt.wantErr {
				t.Errorf("QueryUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			for _, tmp := range res {
				fmt.Println(tmp)
			}
		})
	}
}

func TestQueryUserById(t *testing.T)  {
	Init()
	type argument struct {
		ctx     context.Context
		ID  	int64
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
			ctx: context.Background(),
			ID: 1,
		}, wantErr: false},
		{
			name: "2", args: argument{
			ctx: context.Background(),
			ID: 2,
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := QueryUserById(tt.args.ctx, tt.args.ID)
			if  (err != nil) != tt.wantErr {
				t.Errorf("QueryUserById() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println(res)
		})
	}
}

func TestUpdateFollowAndFollowerCount(t *testing.T)  {
	Init()
	type argument struct {
		ctx     context.Context
		userId  	int64
		followUserId int64
		num int64
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		//用户表（user表）关注者userId关注数增加、被关注者followUserId粉丝增加
		{
			name: "1", args: argument{
			ctx: context.Background(),
			userId: 1,
			followUserId: 2,
			num: 1,
		}, wantErr: false},
		//用户表（user表）关注者userId关注数减少、被关注者followUserId粉丝数减少
		{
			name: "2", args: argument{
			ctx: context.Background(),
			userId: 1,
			followUserId: 2,
			num: -1,
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UpdateFollowAndFollowerCount(tt.args.ctx, tt.args.userId, tt.args.followUserId, tt.args.num)
			if  (err != nil) != tt.wantErr {
				t.Errorf("QueryUserById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}