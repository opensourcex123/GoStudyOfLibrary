package main

import (
	"fmt"
	messagebus "github.com/vardius/message-bus"
	"sync"
)

// Go 每日一库之 message-bus 非常小巧的异步消息通信库,只能在一个进程中使用
func main() {
	queueSize := 100
	bus := messagebus.New(queueSize)

	var wg sync.WaitGroup
	wg.Add(2)

	_ = bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		fmt.Println(v)
	})

	_ = bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		fmt.Println(v)
	})

	bus.Publish("topic", true)
	wg.Wait()
}
