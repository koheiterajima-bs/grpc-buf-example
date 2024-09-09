// Package myipapis は、pb 層
package myipapis

import (
	"context"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/plusmedi/mhv2-backends/pkg/service"

	myipv1 "github.com/takekazuomi/grpczap01/pkg/apis/myip/v1alpha1"
	myiplog "github.com/takekazuomi/grpczap01/pkg/utils/logger"

	"github.com/takekazuomi/grpczap01/pkg/myip"

	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Service implements the MyIP Service API.
type Service struct {
	reporters service.Reporters

	myipv1.UnimplementedMyipServiceServer
}

func New(ctx context.Context) (*Service, error) {
	// logger はコンテキストから取得する
	log := ctxzap.Extract(ctx)
	_ = log

	s := &Service{}

	// レポーター群をs.reportersに登録する
	for _, i := range []struct {
		Name     string
		Reporter service.Reporter
	}{
		{"ipify", service.ReporterFunc(s.IpifyReporter)},
	} {
		if err := s.reporters.AddReporter(i.Name, i.Reporter); err != nil {
			log.Error("AddReporter", zap.String("name", i.Name), zap.Error(err))
			return nil, err
		}
	}

	return s, nil
}

func Register(s grpc.ServiceRegistrar, srv *Service) {
	myipv1.RegisterMyipServiceServer(s, srv)
}

// Create request.
func (s *Service) Create(ctx context.Context, req *myipv1.CreateRequest) (*myipv1.CreateResponse, error) {
	// grpc_zap でctxに設定された、zap logger を context から取得する。
	// もしctxにzap loggerが設定されていなかった場合は、nullLogger が返される。
	log := grpc_zap.Extract(ctx)

	log.Debug("args", myiplog.Any("req", req))

	if err := validateCreateRequest(ctx, req); err != nil {
		return nil, err
	}

	r, err := myip.Create(ctx, myip.IPType(req.IpType))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "ip_type was not found")
	}
	return &myipv1.CreateResponse{
		Myip: &myipv1.MyIP{
			Name:       r.Name,
			IpType:     myipv1.IPType(r.IPType),
			IpAddress:  r.IPAddress,
			State:      myipv1.MyIP_STATE_ACCEPTED,
			CreateTime: timestamppb.New(r.CreateTime),
			UpdateTime: timestamppb.New(r.UpdateTime),
		},
	}, nil
}

// validateCreateRequest CreateRequest のバリデーション
// proto に書いた定義を使って自動生成できるなら、ここは不用
func validateCreateRequest(ctx context.Context, req *myipv1.CreateRequest) error {
	log := grpc_zap.Extract(ctx)

	for _, v := range []struct {
		msg  string
		name string
		test func() bool
	}{
		{"required", "CreateRequest", func() bool { return req != nil }},
		{"required", "ip_type", func() bool { return req.IpType != 0 }},
	} {
		if !v.test() {
			// TODO こんなところで書きたくない。grpcの出口で書くと
			// 出るようにしたい。stack trace 込で
			log.Error("Invalid Argument",
				zap.String("name", v.name),
				zap.String("msg", v.msg),
			)
			return status.Errorf(codes.InvalidArgument, "%q %s", v.name, v.msg)
		}
	}

	return nil
}

// Get status
func (s *Service) Get(ctx context.Context, req *myipv1.GetRequest) (*myipv1.GetResponse, error) {
	log := grpc_zap.Extract(ctx)
	log.Debug("args", myiplog.Any("req", req))

	if err := validateGetRequest(ctx, req); err != nil {
		return nil, err
	}

	r, ok := myip.Get(ctx, req.Name)
	if !ok {
		return nil, status.Error(codes.NotFound, "resource name was not found")
	}
	return &myipv1.GetResponse{
		Myip: &myipv1.MyIP{
			Name:       r.Name,
			IpType:     myipv1.IPType(r.IPType),
			IpAddress:  r.IPAddress,
			State:      myipv1.MyIP_STATE_ACCEPTED,
			CreateTime: timestamppb.New(r.CreateTime),
			UpdateTime: timestamppb.New(r.UpdateTime),
		},
	}, nil
}

func validateGetRequest(ctx context.Context, req *myipv1.GetRequest) error {
	log := grpc_zap.Extract(ctx)

	for _, v := range []struct {
		msg  string
		name string
		test func() bool
	}{
		{"required", "GetRequest", func() bool { return req != nil }},
		{"required", "name", func() bool { return len(req.Name) != 0 }},
	} {
		if !v.test() {
			log.Error("Invalid Argument",
				zap.String("name", v.name),
				zap.String("msg", v.msg),
			)
			return status.Errorf(codes.InvalidArgument, "%q %s", v.name, v.msg)
		}
	}

	return nil
}

// Close はサービスを停止する.
func (s *Service) Close() {
	// TODO: DBコネクションプールの破棄。普段やってない。
	// TODO(takuma): cnsコネクションをクローズする
}

// Register は service.Service を実装する.
func (s *Service) Register(gr grpc.ServiceRegistrar) {
	myipv1.RegisterMyipServiceServer(gr, s)
}

// Reporters は service.Service を実装する.
func (s *Service) Reporters(ctx context.Context) *service.Reporters {
	return &s.reporters
}
