# coreproxy 対応 grpcのツアー

## プライベートモジュール

- https://github.com/plusmedi/go-coreg
- https://github.com/plusmedi/mhv2-backends/tree/main/pkg

## service runner と factory

復数のサービスで、単一のgrpcサーバーを使うために、サービスを`factory pattern`で実装して、grpcサーバーに登録する[service runner](https://github.com/plusmedi/mhv2-backends/tree/main/pkg/service)を使う。サービスのファクトリーにフラグを渡すのも、このあたりでやる。

main.goで下記モジュールをimportし。

```go
"github.com/plusmedi/mhv2-backends/pkg/service"
"github.com/plusmedi/mhv2-backends/pkg/service/runner"
```

適当な名前を付けて、`runner.New("myipapis")` して、サービスを実装したパッケージに、`service.Factory` を実装し、`(* Runner)Register()` する。このあたりは、すこしinterfaceが複雑だが、基本コピペで行ける。

## StatusService

各サービスは、Heath Checkと設定確認のため。StatusServiceを実装する。


### grpc server の Interceptor

mhv2-backends/pkg/service"で作成される、grpc サーバーは、下記のInterceptorを設定している。

```go
s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			reqlog.UnaryServerInterceptor(),
			rpclog.UnaryServerInterceptor(),
			grpc_validator.UnaryServerInterceptor(),
			cache.New(conf.CacheTTL).UnaryServerInterceptor(),
			richerror.UnaryServerInterceptor,
		),
		grpc_middleware.WithStreamServerChain(
			reqlog.StreamServerInterceptor(),
			rpclog.StreamServerInterceptor(),
			grpc_validator.StreamServerInterceptor(),
			richerror.StreamServerInterceptor,
		),
	)
```

参考
- [mhv2-backends/pkg/service/runner/server.go](https://github.com/plusmedi/mhv2-backends/blob/main/pkg/service/runner/server.go#L50-L65)
- [Package cache](https://github.com/plusmedi/mhv2-backends/tree/main/pkg/service/runner/cache)

## logging

zap を使って構造化ロギングを実装する。ACAで動かすと、app -> zap -> stdout -> log analytics と流れて保存される。また、ローカルで`docker run -d`で動かした場合は、`docker logs` でログを見ることができ、`go run `で動かすと、コンソールに出る。

loggerはパッケージのlogger["github.com/plusmedi/go-coreg/zap/logger/v2"](https://github.com/plusmedi/go-coreg/tree/main/zap/logger/v2)と、contextに割り当てられたloggerの["github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"](https://pkg.go.dev/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/logging/zap)の２種類がある。多くの場合は、grpcのリクエストのコンテキストで動くので、go のcontextからloggerを取得して利用する。
例えば、下記のようにgrpc_zap でctxに設定された、zap logger を context から取得する。もしctxにzap loggerが設定されていなかった場合は、nullLogger が返される。

```go
log := grpc_zap.Extract(ctx)
```

## error

[gRPC Richer Error](https://www.grpc.io/docs/guides/error/#richer-error-model) を使う。GoのエラーをRicher error modelに変換するライブラリを用意した。

https://github.com/plusmedi/go-coreg/tree/main/grpc/richerror

richerrorを通して置くと、grpc interceptorで、Go Errorから、richer error model に変換される。

## validation

validationは、[protoc-gen-validate (PGV)](https://github.com/envoyproxy/protoc-gen-validate)を使い。grpc interceptor でチェックされる。

## TODO

- [ ] MySQLの例
  - [ ] sql-migrate
  - [ ] sqlboiler
  - [ ] mysql container
- [ ] test
  - [ ]


## メモ

https://baleen-studio.slack.com/archives/C01GSGBSRCL/p1661990650058419


1.プロジェクトの新規作成
・単体
　・gRPC+MySQLのサービスを実行可能なひな形までのフローと使用するツールやスクリプトなど
　・複数プロジェクトの場合（gRPC x 2、gRPCとCronやQueueトリガーとの組み合わせ）の推奨設定
　・StatusServiceについて
　・request_idについて
・他のサービスとの連携
　・Core Proxyについての役割
　・認証・認可の情報の活用
　・UserやHospitalなど、他のサービスの呼び方
　・ミドルの環境や設定(MySQL、ストレージなど）
2.コード管理
・Githubを使うことはしってますが、プロジェクトをどう分ける？まとめる？とか聞きたいです。
3.テスト
・不明
4.デプロイ
・不明
5.担当
・不明
2,3,4は、今月から実装に初参加する人と同じレベルで、本当に何もわかってない状況です。
5は、1人で開発する前提で進めて良いのか？それとも担当を細分化する予定なのか？そのあたりがわかるともうちょっと効率よく進められるかと。
説明会の開催お待ちしております。
:okサイン:
1


