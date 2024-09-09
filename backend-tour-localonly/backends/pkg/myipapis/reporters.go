package myipapis

import (
	"context"

	"github.com/plusmedi/mhv2-backends/pkg/service"
	"github.com/takekazuomi/grpczap01/pkg/myip"
)

// IpifyReporter は、myip が依存するサービスの状態を報告する
func (s *Service) IpifyReporter(ctx context.Context) *service.Report {
	_, err := myip.GetMyIP(context.Background(), myip.IPV4)
	// TODO もう少しなんとかする
	if err != nil {
		return &service.Report{
			Status: service.StatusUnavailable,
			Text:   err.Error(),
		}
	}
	return &service.Report{
		Status: service.StatusNormal,
		Text:   "正常",
	}
}
