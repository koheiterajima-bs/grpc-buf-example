// Package logger は、zap field の helper、名前は暫定
package logger

import (
	"encoding/json"

	"go.uber.org/zap"
)

// Any は、zap.Any()のJSON展開版
func Any(key string, value interface{}) zap.Field {
	p, err := json.Marshal(value)
	if err == nil {
		return zap.Any(key, p)
	}
	return zap.Any(key, value)
}
