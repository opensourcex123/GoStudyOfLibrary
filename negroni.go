package main

import (
	"fmt"
	"github.com/urfave/negroni"
	"net/http"
	"time"
)

// negroni是一个专注于 HTTP 中间件的库。它小巧，无侵入，鼓励使用标准库net/http的处理器（Handler）
func index(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Fprintf(w, "home page")
	fmt.Printf("index of elasped:%fs\n", time.Since(start).Seconds())
}

func greeting(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	name := r.FormValue("name")
	if name == "" {
		name = "world"
	}

	fmt.Fprintf(w, "hello %s", name)
	fmt.Printf("greeting elasped:%fs\n", time.Since(start).Seconds())
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/greeting", greeting)

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":8000", n)
}
