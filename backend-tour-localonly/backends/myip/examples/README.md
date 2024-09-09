# 使い方

## 起動

ホストで実行

```sh
$ go run ./cmd/myipapis
```

コンテナで実行

```
make build 

```



Createを読んで、リソースの作成

```sh
$ grpcurl -use-reflection -plaintext \
 -d '{"ip_type":"1"}' \
 ${MYIPAPIS_ADDR} \
 myip.v1alpha1.MyipService.Create
{
  "myip": {
    "name": "myips/01gc5wzhhfsj35g9sqhc45t7yv",
    "ipType": "IP_TYPE_IPV4",
    "ipAddress": "116.220.116.42",
    "state": "STATE_ACCEPTED",
    "createTime": "2022-09-05T04:01:49.615804796Z",
    "updateTime": "2022-09-05T04:01:49.615804796Z"
  }
}
```

```sh
$ grpcurl -use-reflection -plaintext \
 -d '{"name":"myips/01gc5wzhhfsj35g9sqhc45t7yv"}' \
 ${MYIPAPIS_ADDR} \
 myip.v1alpha1.MyipService.Get
{
  "myip": {
    "name": "myips/01gc5wzhhfsj35g9sqhc45t7yv",
    "ipType": "IP_TYPE_IPV4",
    "ipAddress": "116.220.116.42",
    "state": "STATE_ACCEPTED",
    "createTime": "2022-09-05T04:01:49.615804796Z",
    "updateTime": "2022-09-05T04:01:49.615804796Z"
  }
}
```

## 以下参考

```sh
$ grpcurl -use-reflection -plaintext ${MYIPAPIS_ADDR} list
grpc.reflection.v1alpha.ServerReflection
myip.v1alpha1.MyipService
status.v1alpha1.StatusService
```

```sh
$ grpcurl -use-reflection -plaintext ${MYIPAPIS_ADDR} list myip.v1alpha1.MyipService
myip.v1alpha1.MyipService.Create
myip.v1alpha1.MyipService.Get
```

```sh
$ grpcurl -use-reflection -plaintext ${MYIPAPIS_ADDR} describe myip.v1alpha1.MyipService.Create
myip.v1alpha1.MyipService.Create is a method:
rpc Create ( .myip.v1alpha1.CreateRequest ) returns ( .myip.v1alpha1.CreateResponse );
```
