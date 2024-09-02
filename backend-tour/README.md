# メモ

## 理解までの手順

- 一通りコードを読み、どこのパートで何をやっているか
  - どういった手順で読み解いていくか？
    - proto→pb.go→main.go以外のGo→main.goかね？(それぞれフリガナ振っていく)
- pb.goの各メソッドがどこで使われているか見る。


- APIを追加する、そのためにはprotoに追記
- 外部のAPI、オプションがいくつかあるので、追加してみる

- myip.go or api.goでgrpcurlの出力を担っている（多分api.go,構造体にstateがあるので）
- 郵便番号APIを使う(https://zipcloud.ibsnet.co.jp/doc/api)

## とりあえず動かす

- ホストで実行

  ```sh
  go run ./cmd/myipapis
  ```

- MYIPAPIS_ADDRの値がセットされていないので、セットする

```sh
export MYIPAPIS_ADDR="localhost:51051"
```

- 設定確認

```sh
echo $MYIPAPIS_ADDR
```

- grpcurlにて呼んで、リソースの作成

```sh
grpcurl -use-reflection -plaintext \
-d '{"ip_type":"1"}' \
${MYIPAPIS_ADDR} \
myip.v1alpha1.MyipService.Create
```

```sh
grpcurl -use-reflection -plaintext \
 -d '{"name": "myips/ここにuuidが入る"}' \
 ${MYIPAPIS_ADDR} \
 myip.v1alpha1.MyipService.Get
```

```sh
grpcurl -use-reflection -plaintext \
 -d '{"name": "myips/01j6rpsq46qq0rs649qsq3jqhz"}' \
 ${MYIPAPIS_ADDR} \
 myip.v1alpha1.MyipService.Delete
```

## どこのパートで何をやっているか

- ディレクトリ構造
  - api-specs：API定義を保存するリポジトリ
  - backends/myip：Backendのコードを保存するリポジトリ

- api-specsディレクトリ
  - myip.proto：自分のAPIを取得するためのipify APIというAPIを使っており、データ構造やそれに関連する列挙型、フィールド、タイムスタンプなどを定義するために使われる
  - api.proto：myip.protoをインポートしている、gRPC APIのインターフェース、つまりサーバーとクライアントがどのように通信するかに関する定義に使われる

.
└── api-specs
    ├── myipapis
    │   ├── myip/v1alpha1
    │   │   ├── api.proto
    │   │   └── myip.proto
    │   ├── buf.lock
    │   └── buf.yaml
    └── buf.work.yaml

- backends/myipディレクトリ

.
└── backends/myip
    ├── cmd/myipapis
    │   └── main.go (myipapis実行ファイル)
    ├── deployments
    │   ├── logs
    │   │   └── .gitignore
    │   ├── modules
    │   │   └── aca.bicep
    │   ├── .envrc
    │   ├── env.json
    │   ├── main.bicep
    │   ├── Makefile
    │   └── README.md
    ├── examples
    │   └── README.md
    ├── internal/pkg/poc1
    │   ├── .envrc
    │   ├── gitignore
    │   ├── buf.gen.yaml
    │   ├── Makefile
    │   └── README.md
    ├── pkg
    │   ├── apis
    │   │   ├── myip/v1alpha1
    │   │   │   ├── api_grpc.pb.go
    │   │   │   ├── api.pb.go
    │   │   │   └── myip.pb.go
    │   │   └── staticcheck.conf
    │   ├── myip
    │   │   ├── ipifyclient_test.go
    │   │   ├── ipifyclient.go
    │   │   ├── myip_test.go
    │   │   └── myip.go
    │   ├── myipapis
    │   │   ├── factory
    │   │   │   ├── factory_test.go
    │   │   │   └── factory.go
    │   │   ├── api.go
    │   │   └── reporters.go
    │   └── utils/logger
    │       └── field.go
    ├── scripts
    │   └── install-deps.sh
    ├── .editorconfig
    ├── .envrc
    ├── gitignore
    ├── buf.gen.yaml
    ├── go.mod
    ├── go.sum
    ├── go.work
    ├── go.work.sum
    ├── Makefile
    └── staticcheck.conf


#backends/myip
##cmd/myipapis
###main.go
##deployments
###logs
####.gitignore
###modules
####aca.bicep
###.envrc
###env.json
###main.bicep
###Makefile
###README.md
##examples
###README.md
##internal/pkg/poc1
###.envrc
###gitignore
###buf.gen.yaml
###Makefile
###README.md
##pkg
###apis
####myip/v1alpha1
#####api_grpc.pb.go
#####api.pb.go
#####myip.pb.go
####staticcheck.conf
###myip
####ipifyclient_test.go
####ipifyclient.go
####myip_test.go
####myip.go
###myipapis
####factory
#####factory_test.go
#####factory.go
####api.go
####reporters.go
###utils/logger
####field.go
##scripts
###install-deps.sh
##.editorconfig
##.envrc
##gitignore
##buf.gen.yaml
##go.mod
##go.sum
##go.work
##go.work.sum
##Makefile
##staticcheck.conf

(疑問)
- goファイルで、api_grpc.pb.go、api.pb.go、myip.pb.goらのメソッドは全く使われていない？
- main.goで使われている関数、メソッドはほぼrunner.goのもの？
- go runででてくるコマンドは、どこから出力されている？
- よくわからんが、Registerで登録し、Runで実行する
- zapについてみてみる->実行結果はzapで構造化したものとはわかったが、、

## 知らなかったこと

- enum(列挙型)は、ProtocolBufferファイルの中で使われるデータ型で、事前定義された数値のセットを表現するために使われる。意味のある名前を持つ整数のリストとして扱われる。

```proto
enum IPType {
	IP_TYPE_UNSPECIFIED = 0;
	IP_TYPE_IPV4 = 1;
	IP_TYPE_UNIVERSAL = 2;
}
```

- grpc-zapとは？
  - gRPCサーバーやクライアントで発生するリクエストやレスポンス、エラーを高速かつ効率的にロギングするためにzapロガーを統合するライブラリ
  - gRPCサービスのデバッグ、モニタリング、パフォーマンス追跡に役立つ

## 疑問点

- `backends/myip/pkg/apis/myip/v1alpha1/`以下に各pbファイル格納されているが、
  - buf.gen.yamlファイルを見ても、`backends/myip/pkg/apis/myip/v1alpha1/`以下からコンパイルした形式が見られない？


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


