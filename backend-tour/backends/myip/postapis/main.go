// クライアントの実装
package main

/*
// これはHTTPクライアントを使ってREST APIにアクセスしている


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
*/

/*

// gRPCとREST APIを混同してしまっている
// grpc.NewClientはgRPCサーバーと通信するためのクライアントを作成するもの
// gRPCサーバーについてよくわかっていないのでは、REST APIとgRPCについて、もう一度何が違うのか、学んでみる
// 外部で公開されているAPIでgRPCを利用しているものを探す

func main() {
	conn, err := grpc.NewClient("https://zipcloud.ibsnet.co.jp/api/search?zipcode=1600008", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPostAnswerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayAddress(ctx, &pb.PostRequest{Requestpostcode: 1600008})
	if err != nil {
		log.Fatalf("郵便番号が不正: %v", err)
	}

	for _, post := range r.GetResults() {
		log.Printf("Address1: %s", post.GetAddress1())
		log.Printf("Address2: %s", post.GetAddress2())
		log.Printf("Address3: %s", post.GetAddress3())
	}
}
*/
