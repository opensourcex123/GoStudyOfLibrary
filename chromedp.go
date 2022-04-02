package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

type InviteApplication struct {
	CompanyUrl      string
	CompanyImage    string
	CompanyName     string
	RecruitTime     string
	CompanyAct      string
	CompanyActTime  string
	DeliveryLink    string
	WrittenTestTime string
	InterviewTime   string
	OfferTime       string
}

func main() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	url := "https://nowpick.nowcoder.com/w/school/schedule?property=4715"
	sel := "//div[1]/div[1]/div[2]/div/div[3]/div[3]/ul/li[@class='list-item']"
	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("body"),
		chromedp.Sleep(time.Second),
		chromedp.Nodes(sel, &nodes, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}
	var inviteApplication InviteApplication
	var attr map[string]string
	for i := 0; i < len(nodes); i++ {
		// 获取公司url
		chromedp.Run(ctx,
			chromedp.Attributes(fmt.Sprintf("//..//li[%d]/a", i+1), &attr, chromedp.NodeVisible),
		)
		inviteApplication.CompanyUrl = "https://nowpick.nowcoder.com/" + attr["href"]
		// 获取公司图片地址
		chromedp.Run(ctx,
			chromedp.Attributes(fmt.Sprintf("//..//li[%d]/a/div[1]/div/img", i+1), &attr, chromedp.NodeVisible),
		)
		inviteApplication.CompanyImage = attr["src"]
		// 获取公司名称
		chromedp.Run(ctx,
			chromedp.Text(fmt.Sprintf("//..//li[%d]/a/div[1]/div/div[2]", i+1), &inviteApplication.CompanyName, chromedp.NodeVisible),
		)
		// 获取内推时间
		chromedp.Run(ctx,
			chromedp.Text(fmt.Sprintf("//..//li[%d]/a/div[2]/div[2]/span", i+1), &inviteApplication.RecruitTime, chromedp.NodeVisible),
		)
		log.Printf("%#v", inviteApplication)
	}
}
