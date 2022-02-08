package main

import (
	"fmt"
	"github.com/uniplaces/carbon"
	"time"
)

// carbon库是一个时间扩展库，基于 PHP 的carbon库编写。提供易于使用的接口
func main() {
	fmt.Printf("Right now is %s\n", carbon.Now().DateTimeString())
	fmt.Println(carbon.Now().EndOfDay())

	today, _ := carbon.NowInLocation("Japan")
	fmt.Printf("Right now in Japan is %s\n", today)

	fmt.Printf("Tomorrow is %s\n", carbon.Now().AddDay())
	fmt.Printf("Last week is %s\n", carbon.Now().SubWeek())

	nextOlympics, _ := carbon.CreateFromDate(2020, time.August, 5, "Europe/London")
	nextOlympics = nextOlympics.AddYears(4)
	fmt.Printf("Next olympics are in %d\n", nextOlympics.Year())

	if carbon.Now().IsWeekend() {
		fmt.Println("Happy Weekend!")
	}
}
