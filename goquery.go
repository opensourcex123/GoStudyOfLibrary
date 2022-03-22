package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

// goquery是用 Go 语言编写的一个类似于 jQuery 的库。
// 它基于 HTML 解析库net/html和 CSS 库cascadia，提供与 jQuery 相近的接口

func BaiduHotSearch() {
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".s-hotsearch-content .hotsearch-item").Each(func(i int, selection *goquery.Selection) {
		content := selection.Find(".title-content-title").Text()
		fmt.Printf("%d: %s\n", i, content)
	})
}

func main() {
	BaiduHotSearch()
}
