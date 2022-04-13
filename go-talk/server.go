package main

import (
	"github.com/rsms/gotalk"
	"log"
)

// gotalk专注于进程间的通信，致力于简化通信协议和流程
func main() {
	gotalk.Handle("echo", func(in string) (string, error) {
		return in, nil
	})
	if err := gotalk.Serve("tcp", ":8080", nil); err != nil {
		log.Fatal(err)
	}
}
