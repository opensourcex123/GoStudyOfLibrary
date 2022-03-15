package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

// cron一个用于管理定时任务的库，用 Go 实现 Linux 中crontab这个命令的效果
func main() {
	c := cron.New()

	c.AddFunc("@every 1s", func() {
		fmt.Println("tick every 1 second")
	})

	c.Start()
	time.Sleep(time.Second * 5)
}
