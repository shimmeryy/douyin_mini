package db

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestQueryByTime(t *testing.T) {
	Init()
	type args struct {
		ctx       context.Context
		last_time int64
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
				ctx:       context.Background(),
				last_time: time.Now().Unix(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryByTime(tt.args.ctx, tt.args.last_time)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryByTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i := range got {
				fmt.Printf("%+v\n", got[i])
			}
		})
	}
}
