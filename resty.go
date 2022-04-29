package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
)

// resty是 Go 语言的一个 HTTP client 库。resty功能强大，特性丰富
func main() {
	client := resty.New()

	resp, err := client.R().Get("https://baidu.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response Info:")
	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Status:", resp.Status())
	fmt.Println("Proto:", resp.Proto())
	fmt.Println("Time:", resp.Time())
	fmt.Println("Received At:", resp.ReceivedAt())
	fmt.Println("Size:", resp.Size())
	fmt.Println("Headers:")
	for key, val := range resp.Header() {
		fmt.Println(key, "=", val)
	}
	for i, cookie := range resp.Cookies() {
		fmt.Printf("cookie%d: name:%s value:%s\n", i, cookie.Name, cookie.Value)
	}
}
