package main

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v2"
)

// gentleman是一个功能齐全、插件驱动的 HTTP 客户端。
// gentleman以扩展性为原则，可以基于内置的或第三方插件创建具有丰富特性的、可复用的 HTTP 客户端
func main() {
	cli := gentleman.New() // 此对象可复用
	cli.URL("https://dog.ceo")

	req := cli.Request()
	req.Path("/api/breeds/image/random")
	req.SetHeader("Client", "gentleman")

	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request error %v\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %dn", res.StatusCode)
		return
	}

	fmt.Printf("Body: %s", res.String())
}
