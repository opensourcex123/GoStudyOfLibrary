package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/guonaihong/gout"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type DiscussPosts struct {
	PostId    int    `json:"postId"`
	PostTitle string `json:"postTitle"`
}
type ArticleData struct {
	DiscussPosts []DiscussPosts `json:"discussPosts"`
	TotalPage    int            `json:"total_page"`
}
type RespArticle struct {
	Data ArticleData `json:"data"`
}

func main() {
	var resp RespArticle
	var filter = []string{"?", "？", "聘", "牛客"} // 过滤字数组
	// 获取随机tag_id和页数
	tagId := []int{639, 640, 641, 642, 644, 645, 680, 733, 894, 4396, 1194}
	rand.Seed(time.Now().UnixNano())
	// 获取该tag下的总页数
	tagIndex := rand.Intn(11)
	titleUrl := fmt.Sprintf("https://www.nowcoder.com/discuss/experience/json?tagId=%d&page=1", tagId[tagIndex])
	gout.GET(titleUrl).
		BindJSON(&resp).Do()
	totalPage := resp.Data.TotalPage
	// 随机爬取一个tag下的3篇文章
	for i := 0; i < 1; i++ {
		var title string
		var text string
		var content string
		// 获取完全随机的url
		page := rand.Intn(totalPage + 1)
		titleUrl = fmt.Sprintf("https://www.nowcoder.com/discuss/experience/json?tagId=%d&page=%d", tagId[tagIndex], page)
		// 获取文章标题
		gout.GET(titleUrl).
			BindJSON(&resp).Do()
		articleIndex := rand.Intn(len(resp.Data.DiscussPosts))
		title = resp.Data.DiscussPosts[articleIndex].PostTitle
		log.Println("title:", title)
		// 排除问答文章及招聘广告
		isLow := false
		for _, item := range filter {
			if strings.Contains(resp.Data.DiscussPosts[articleIndex].PostTitle, item) {
				isLow = true
				break
			}
		}
		if isLow {
			continue
		}
		// 获取文章内容
		c := colly.NewCollector()
		contentUrl := fmt.Sprintf("https://www.nowcoder.com/discuss/%d", 923146)
		c.OnHTML("body .post-topic-main .post-topic-des", func(e *colly.HTMLElement) {
			text = strings.TrimSpace(e.Text)
			//log.Println("内容：", text)
		})
		c.OnResponse(func(r *colly.Response) {
			contentRaw := string(r.Body)
			contentRaw = strings.TrimSpace(strings.Split(contentRaw, "nc-post-content\">")[1])
			contentRaw = strings.TrimSpace(strings.Split(contentRaw, "<div class=\"clearfix\">")[0])
			content = contentRaw
			reg := regexp.MustCompile("<pre")
			content = reg.ReplaceAllString(content, "<pre style=\"white-space:pre-wrap; word-wrap:break-word; word-break:break-all\"")
			reg = regexp.MustCompile("<img")
			content = reg.ReplaceAllString(content, "<img width=100%")
			//log.Println(content)
		})
		c.Visit(contentUrl)
		// 排除低质量文章
		if len(text) < 200 {
			//log.Println("低质量文章")
			continue
		}
		if strings.Contains(text, "牛客") {
			log.Println("不合理文章")
			continue
		}
	}
}