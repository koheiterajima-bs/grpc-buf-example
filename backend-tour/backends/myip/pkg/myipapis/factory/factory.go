// Package factory は github.com/plusmedi/coreg-sandbox/takekazuomi/backend-tour/backends/myip/pkg/myip のファクトリー
package factory

import (
	"context"

	"github.com/plusmedi/mhv2-backends/pkg/service"
	"github.com/spf13/pflag"
	"github.com/takekazuomi/grpczap01/pkg/myipapis"
)

type Factory struct {
	// フラグを追加する場合は、ここにメンバーを用意する
}

func New() *Factory {
	return &Factory{}
}

func (f *Factory) Flags() *pflag.FlagSet {
	fset := pflag.NewFlagSet("myipapis", pflag.ExitOnError)
	// フラグを追加する場合は、ここに設定を書く。
	return fset
}

func (f *Factory) New(ctx context.Context) (service.Service, error) {
	// フラグを追加する場合は、ここで渡す。
	s, err := myipapis.New(ctx)
	return s, err
}
