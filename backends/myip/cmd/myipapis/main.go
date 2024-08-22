package main

import (
	"os"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/takekazuomi/grpczap01/pkg/myipapis/factory" // これが何をするリポジトリか？
	"go.uber.org/zap"

	"github.com/plusmedi/go-coreg/zap/logger/v2"
	"github.com/plusmedi/mhv2-backends/pkg/service"
	"github.com/plusmedi/mhv2-backends/pkg/service/runner"
)

// ログ出力
var log = logger.L()

func main() {
	if run() != nil {
		os.Exit(1)
	}
}

func run() error {
	// 関数の最後に実行、無名関数で戻り値は不要
	// バッファリングされているログデータを全て書き出し、ロガーが保持する全てのリソースを適切にクリーンアップする。
	defer func() { _ = log.Sync() }()

	// New関数にて引数に"myipapis"を入れ、Runnerを返す(よくわからない、、)
	r := runner.New("myipapis")

	// for rangeで
	for _, sf := range []struct {
		Name    string
		Factory service.Factory
	}{
		{"myip", factory.New()},
	} { // 何かしらに登録する？
		if err := r.Register(sf.Name, sf.Factory); err != nil {
			log.Error("r.Register", zap.Error(err), zap.String("name", sf.Name))
			os.Exit(1)
		}
	}

	// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
	grpc_zap.ReplaceGrpcLoggerV2(logger.Get("google.golang.org/grpc/grpclog"))

	// Runの引数に何も入っていない場合のエラーハンドリング
	err := r.Run(nil)
	if err != nil {
		log.Debug("r.Run", zap.Error(err))
	}
	return err
}
