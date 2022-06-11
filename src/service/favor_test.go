package service

import (
	"context"
	"fmt"
	"testing"
	"tiktok/src/dal"
	"tiktok/src/handlers"
)

func TestOperateFavor(t *testing.T) {
	dal.Init()
	type argument struct {
		ctx context.Context
		req handlers.FavorOperateParam
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx: context.Background(),
				req: handlers.FavorOperateParam{
					UserId:     1,
					VideoId:    24,
					ActionType: 1,
				},
			}, wantErr: false},
		{
			name: "2", args: argument{
				ctx: context.Background(),
				req: handlers.FavorOperateParam{
					UserId:     1,
					VideoId:    10,
					ActionType: 2,
				},
			}, wantErr: false},
		{
			name: "3", args: argument{
				ctx: context.Background(),
				req: handlers.FavorOperateParam{
					UserId:     999999,
					VideoId:    24,
					ActionType: 1,
				},
			}, wantErr: true},
		{
			name: "4", args: argument{
				ctx: context.Background(),
				req: handlers.FavorOperateParam{
					UserId:     999999,
					VideoId:    10,
					ActionType: 2,
				},
			}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FavorServiceInstance().OperateFavor(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("OperateFavor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueryFavorByUserId(t *testing.T) {
	dal.Init()
	type argument struct {
		ctx context.Context
		req handlers.FavorQueryParam
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx: context.Background(),
				req: handlers.FavorQueryParam{
					UserId: 1,
				},
			}, wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := FavorServiceInstance().QueryFavorByUserId(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryFavorByUserId() error = %v, wantErr = %v", err, tt.wantErr)
			}
			for _, tmp := range res {
				fmt.Println(tmp)
			}
		})
	}
}

func TestQueryFavorVideoByUserId(t *testing.T) {
	dal.Init()
	type argument struct {
		ctx context.Context
		req handlers.FavorQueryParam
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx: context.Background(),
				req: handlers.FavorQueryParam{
					UserId: 1,
				},
			}, wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := FavorServiceInstance().QueryFavorVideoByUserId(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryFavorVideoByUserId() error = %v, wantErr = %v", err, tt.wantErr)
			}
			for _, tmp := range res {
				fmt.Println(tmp)
			}
		})
	}
}

func TestCountFavorByVideoId(t *testing.T) {
	dal.Init()
	type argument struct {
		ctx context.Context
		req handlers.FavorCountParam
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx: context.Background(),
				req: handlers.FavorCountParam{
					VideoId: 15,
				},
			}, wantErr: false,
		},
		{
			name: "2", args: argument{
				ctx: context.Background(),
				req: handlers.FavorCountParam{
					VideoId: 999,
				},
			}, wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := FavorServiceInstance().CountFavorByVideoId(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountFavorByVideoId() error = %v, wantErr = %v", err, tt.wantErr)
			}
			fmt.Println(res)
		})
	}
}

func TestCheckIsFavored(t *testing.T) {
	dal.Init()
	type argument struct {
		ctx context.Context
		req handlers.FavorCheckParam
	}
	tests := []struct {
		name    string
		args    argument
		wantErr bool
	}{
		{
			name: "1", args: argument{
				ctx: context.Background(),
				req: handlers.FavorCheckParam{
					UserId:  1,
					VideoId: 15,
				},
			}, wantErr: false,
		},
		{
			name: "2", args: argument{
				ctx: context.Background(),
				req: handlers.FavorCheckParam{
					UserId:  10,
					VideoId: 999,
				},
			}, wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := FavorServiceInstance().CheckIsFavored(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckIsFavored() error = %v, wantErr = %v", err, tt.wantErr)
			}
			fmt.Println(res)
		})
	}
}
