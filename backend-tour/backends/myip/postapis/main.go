// クライアントの実装
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// APIのレスポンス全体を格納するための構造体
type ApiResponse struct {
	Message string `json:"message"`
	Results []Post `json:"results"`
	Status  int    `json:"status"`
}

// APIのレスポンスデータを格納するための構造体
type Post struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Address3 string `json:"address3"`
	Prefcode string `json:"prefcode"`
	Zipcode  string `json:"zipcode"`
}

func main() {
	// APIのエンドポイント
	url := "https://zipcloud.ibsnet.co.jp/api/search?zipcode=1600008"

	// リクエストの送信
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("API接続失敗: %v", err)
	}
	defer resp.Body.Close()

	// レスポンスボディの読み取り
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("レスポンスボディの読み取り失敗: %v", err)
	}

	// JSONレスポンスをGoの構造体に変換
	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Fatalf("JSONレスポンスをGoの構造体に変換失敗: %v", err)
	}

	// データを表示
	for _, post := range apiResponse.Results {
		fmt.Println(post.Address1, post.Address2, post.Address3, post.Prefcode, post.Zipcode)
	}
}

// gRPCで作ってみる！！！！！
