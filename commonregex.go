package main

import (
	"fmt"
	cregex "github.com/mingrammer/commonregex"
)

// commonregex库来了，它内置很多常用的正则表达式，开箱即用
func main() {
	text := `John, please get that article on www.linkedin.com to me by 5:00PM on Jan 9th 2012. 4:00 would be ideal, actually. If you have any questions, You can reach me at (519)-236-2723x341 or get in touch with my associate at harold.smith@gmail.com`

	dateList := cregex.Date(text)
	timeList := cregex.Time(text)
	linkList := cregex.Links(text)
	phoneList := cregex.PhonesWithExts(text)
	emailList := cregex.Emails(text)

	fmt.Println("data list:", dateList)
	fmt.Println("time list:", timeList)
	fmt.Println("link list:", linkList)
	fmt.Println("phone list:", phoneList)
	fmt.Println("email list:", emailList)
}
