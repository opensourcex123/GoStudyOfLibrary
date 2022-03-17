package main

import (
	"fmt"
	"github.com/bamzi/jobrunner"
	"time"
)

// jobrunner就是其中一个用来执行异步任务的 Go 语言库。
// 得益于强大的cron库，再搭配jobrunner的任务状态监控，jobrunner非常易于使用

type GreetingJob struct {
	Name string
}

func (g GreetingJob) Run() {
	fmt.Println("Hello ", g.Name)
}

func main() {
	jobrunner.Start()
	jobrunner.Schedule("@every 3s", GreetingJob{Name: "dj"})

	time.Sleep(time.Second * 6)
}
