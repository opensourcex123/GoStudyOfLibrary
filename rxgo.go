package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

// ReactiveX，简称为 Rx，是一个异步编程的 API
// RxGo 是 Rx 的 Go 语言实现。借助于 Go 语言简洁的语法和强大的并发支持（goroutine、channel），Rx 与 Go 语言的结合非常完美。
func main() {
	observable := rxgo.Just(1, 2, 3, 4, 5)()
	ch := observable.Observe()
	for item := range ch {
		if item.Error() {
			fmt.Println("error:", item.E)
		}
		fmt.Println(item.V)
	}
}
