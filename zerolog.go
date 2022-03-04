package main

import "github.com/rs/zerolog/log"

// zerolog相比zap更进了一步，它的 API 设计非常注重开发体验和性能。zerolog只专注于记录 JSON 格式的日志，号称 0 内存分配
func main() {
	log.Print("hello world") // 类似官方log库，但输出的是json格式
}
