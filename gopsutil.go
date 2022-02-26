package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

// Go 每日一库之 gopsutil
// gopsutil是 Python 工具库psutil 的 Golang 移植版，可以帮助我们方便地获取各种系统和硬件信息。
// gopsutil为我们屏蔽了各个系统之间的差异，具有非常强悍的可移植性
func main() {
	// 内存使用情况
	v, _ := mem.VirtualMemory()
	fmt.Printf("Total: %v, Available: %v, UsedPercent:%f%%\n", v.Total, v.Available, v.UsedPercent)
	//fmt.Println(v)

	// cpu使用情况
	physicalCnt, _ := cpu.Counts(false)
	logicalCnt, _ := cpu.Counts(true)
	fmt.Printf("physical count:%d logical count:%d\n", physicalCnt, logicalCnt)
}
