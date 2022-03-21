package main

import (
	"fmt"
	"log"
	"net/http"
)

// air是 Go 语言的热加载工具，它可以监听文件或目录的变化，自动编译，重启程序。大大提高开发期的工作效率
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello nfx")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	server := &http.Server{
		Handler: mux,
		Addr:    ":3000",
	}

	log.Fatal(server.ListenAndServe())
}
