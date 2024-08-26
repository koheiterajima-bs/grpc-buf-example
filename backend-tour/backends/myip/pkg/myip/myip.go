// Package myip は、リソースの実装
package myip

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

// MyIP のリソース定義
// TODO: protobufの定義は通信経路のモデルで、sql-boilerの定義はdbのモデル、
// ビジネスロジックのモデルは別にあるとなると結局モデルが３つ必要でヤレヤレと思う。
type MyIP struct {
	Name string

	IPType IPType

	IPAddress string

	CreateTime time.Time

	UpdateTime time.Time

	// TODO: 論理削除では、DeleteTime を入れる。apiでは、論理削除時にはdelete_time は、空になる。
	// time.Timeにコンバートした時に、何を入れるかを決める。
	// github.com/volatiletech/null/v8 を使うと、ビジネスロジック層がDB層に侵食される。
}

var store map[string]*MyIP = make(map[string]*MyIP)

func generateResourceName() string {
	return fmt.Sprintf("myips/%s", strings.ToLower(ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()).String()))
}

// Create は、myip リソースを作成する
func Create(ctx context.Context, ipType IPType) (*MyIP, error) {
	r, err := GetMyIP(ctx, ipType)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	myip := &MyIP{
		Name:       generateResourceName(),
		IPType:     ipType,
		IPAddress:  r.IP,
		CreateTime: now,
		UpdateTime: now,
	}
	store[myip.Name] = myip

	return myip, nil
}

func Get(ctx context.Context, name string) (*MyIP, bool) {

	v, ok := store[name]

	return v, ok
}
