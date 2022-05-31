package ossUtil

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"testing"
)

func TestInitBucket(t *testing.T) {
	tests := []struct {
		name string
		want *oss.Bucket
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitBucket()
		})
	}
}
