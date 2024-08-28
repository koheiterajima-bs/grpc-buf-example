// クライアントの実装
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

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
	var posts Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Fatalf("JSONレスポンスをGoの構造体に変換失敗: %v", err)
	}

	// データを表示
	fmt.Println(posts.Address1, posts.Address2, posts.Address3, posts.Prefcode, posts.Zipcode)
	// for _, post := range posts {
	// 	fmt.Println(post.Address1, post.Address2, post.Address3, post.Prefcode, post.Zipcode)
	// }
}
