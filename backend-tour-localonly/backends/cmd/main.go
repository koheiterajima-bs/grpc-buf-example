package main

import (
	"os"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/takekazuomi/grpczap01/pkg/myipapis/factory"
	"go.uber.org/zap"

	"github.com/plusmedi/go-coreg/zap/logger/v2"
	"github.com/plusmedi/mhv2-backends/pkg/service"
	"github.com/plusmedi/mhv2-backends/pkg/service/runner"
)

var log = logger.L()

func main() {
	if run() != nil {
		os.Exit(1)
	}
}

func run() error {
	defer func() { _ = log.Sync() }()

	r := runner.New("myipapis")

	for _, sf := range []struct {
		Name    string
		Factory service.Factory
	}{
		{"myip", factory.New()},
	} {
		if err := r.Register(sf.Name, sf.Factory); err != nil {
			log.Error("r.Register", zap.Error(err), zap.String("name", sf.Name))
			os.Exit(1)
		}
	}

	// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
	grpc_zap.ReplaceGrpcLoggerV2(logger.Get("google.golang.org/grpc/grpclog"))

	err := r.Run(nil)
	if err != nil {
		log.Debug("r.Run", zap.Error(err))
	}
	return err
}
