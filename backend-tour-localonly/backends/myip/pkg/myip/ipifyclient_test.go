package myip

import (
	"context"
	"reflect"
	"testing"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

func loggedContext() context.Context {
	logger, _ := zap.NewDevelopment()

	return ctxzap.ToContext(context.Background(), logger)
}

func TestGetMyIP(t *testing.T) {
	type args struct {
		ctx context.Context
		t   IPType
	}
	tests := []struct {
		name    string
		args    args
		want    *IpifyResult
		wantErr bool
	}{
		{
			name: "ipv4",
			args: args{
				ctx: loggedContext(),
				t:   1,
			},
			want: &IpifyResult{
				IP: "116.220.116.42",
			},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				ctx: loggedContext(),
				t:   3,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMyIP(tt.args.ctx, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMyIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMyIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
