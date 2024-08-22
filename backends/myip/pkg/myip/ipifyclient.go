package myip

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"strings"

	"go.uber.org/zap"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
)

type IPType int32

const (
	Unspecified IPType = iota
	IPV4
	Universal
	Error404
)

// MyIPError はステータスコードが200以外の場合のエラーを扱う構造体
type MyIPError struct {
	StatusCode int
	URL        string
	Body       string
}

func (m *MyIPError) Error() string {
	b := strings.ReplaceAll(m.Body, "\r", " ")
	b = strings.ReplaceAll(b, "\n", "")
	return fmt.Sprintf("http status code = %d, url = %s, body = %q", m.StatusCode, m.URL, b)
}

const ipifyOrgIPv4 = "https://api.ipify.org?format=json"
const ipifyOrgUniversal = "https://api64.ipify.org?format=json"
const ipifyError404 = "https://api64.ipify.org/format=json"

type IpifyResult struct {
	IP string `json:"ip,omitempty"`
}

var ErrIPTypeInvalid = errors.New("IPType 不正。IPType.IPV4、IPType.Universalが必要です")

func logHTTPResponse(log *zap.Logger, resp *http.Response) {
	// 呼び出し側のLocationが出るように調整
	log = log.WithOptions(zap.AddCallerSkip(1))

	if ce := log.Check(zap.DebugLevel, "trace"); ce != nil {
		dump, err := httputil.DumpResponse(resp, true)
		ce.Write(
			zap.String("resp", string(dump)),
			zap.Error(err),
		)
	}
}

// GetMyIP は、ipify を呼び出して、outbound ip を取得する
// https://www.ipify.org/
func GetMyIP(ctx context.Context, t IPType) (*IpifyResult, error) {
	log := grpc_zap.Extract(ctx)

	var url string
	switch t {
	case IPV4:
		url = ipifyOrgIPv4
	case Universal:
		url = ipifyOrgUniversal
	case Error404:
		url = ipifyError404
	default:
		return nil, ErrIPTypeInvalid
	}

	resp, err := http.Get(url)

	logHTTPResponse(log, resp)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		return nil, &MyIPError{
			StatusCode: resp.StatusCode,
			URL:        url,
			Body:       string(b),
		}
	}
	result := &IpifyResult{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
