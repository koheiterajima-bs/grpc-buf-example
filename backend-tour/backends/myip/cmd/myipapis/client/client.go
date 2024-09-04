package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// 生成されたMyipServiceのGoコードをインポート
	pb "github.com/takekazuomi/grpczap01/pkg/apis/myip/v1alpha1"
)

func main() {
	// サーバーアドレスを設定 (MYIPAPIS_ADDR の値)
	serverAddr := "localhost:51051" // 環境変数 MYIPAPIS_ADDR に対応

	// gRPC接続の設定
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// クライアントの作成
	client := pb.NewMyipServiceClient(conn)

	// リクエストデータの作成 (ip_type フィールドの設定)
	req := &pb.CreateRequest{
		IpType: 1, // grpcurlコマンドで指定された ip_type フィールドの値
	}

	// コンテキストの作成 (タイムアウトを設定)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Createメソッドを呼び出し、レスポンスを取得
	resCreate, err := client.Create(ctx, req)
	if err != nil {
		log.Fatalf("Error calling Create: %v", err)
	}

	resGet, err := client.Get(ctx, &pb.GetRequest{Name: resCreate.Myip.Name})
	if err != nil {
		log.Fatalf("Error calling Get: %v", err)
	}

	// レスポンスの表示
	fmt.Printf("CreateResponse: %v\n", resCreate)
	fmt.Printf("GetResponse: %v\n", resGet)
}
