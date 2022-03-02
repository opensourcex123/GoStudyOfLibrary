package main

import (
	"github.com/maxence-charriere/go-app/v6/pkg/app"
	"log"
	"net/http"
)

func main() {
	h := &app.Handler{
		Title:  "Go-App",
		Author: "Nfx",
	}

	if err := http.ListenAndServe(":8080", h); err != nil {
		log.Fatal(err)
	}
}
