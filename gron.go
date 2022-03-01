package main

import (
	"fmt"
	"github.com/roylee0704/gron"
	"sync"
	"time"
)

// Go 每日一库之 gron
// gron是一个比较小巧、灵活的定时任务库，可以执行定时的、周期性的任务。gron提供简洁的、并发安全的接口
func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	c := gron.New()
	c.AddFunc(gron.Every(5*time.Second), func() {
		fmt.Println("runs every 5 seconds")
	})
	c.Start()

	wg.Wait()
}
