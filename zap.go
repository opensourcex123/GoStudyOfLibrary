package main

import (
	"go.uber.org/zap"
	"time"
)

// uber开源的日志库zap，对性能和内存分配做了极致的优化
func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	url := "http://example.org/api"
	logger.Info("failed to fetch url",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	sugar := logger.Sugar()
	sugar.Infow("failed to fetch url",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("failed to fetch url: %s", url)
}
