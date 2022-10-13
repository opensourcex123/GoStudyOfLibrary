package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"math/rand"
	"time"
)

func GenerateAvatar() string {
	var avatarList []string
	rand.Seed(time.Now().Unix())
	index := rand.Intn(100)*20 + 2
	//log.Println(index)
	url := fmt.Sprintf("http://www.woyaogexing.com/touxiang/index_%d.html", index)

	for len(avatarList) == 0 {
		c := colly.NewCollector()
		c.OnHTML(fmt.Sprintf("#main > div.list-main.mt10.cl > div.list-left.z > div.pMain > div:nth-child(%d) > a[href]", rand.Intn(24)+1), func(e *colly.HTMLElement) {
			link := e.Attr("href")
			//log.Println(link)
			c.Visit(e.Request.AbsoluteURL(link))
		})
		c.OnResponse(func(r *colly.Response) {
			c.OnHTML("#main > div.contMain.mt10 > div.contLeft.z > div.contLeftA > ul > li", func(e *colly.HTMLElement) {
				e.ForEach("img", func(i int, e *colly.HTMLElement) {
					avatarList = append(avatarList, e.Attr("src"))
				})
			})
		})
		c.Visit(url)
	}

	//log.Println("avatar:", avatarList)
	log.Println(len(avatarList))

	return avatarList[rand.Intn(len(avatarList))]
	//log.Println(avatar)
}

func main() {
	for i := 0; i < 10; i++ {
		avatar := GenerateAvatar()
		log.Printf("%d %s", i, avatar)
	}
}
