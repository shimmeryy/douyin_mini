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