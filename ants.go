package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"math/rand"
	"sync"
)

// ants就是其中一个实现 goroutine 池的库
// 一般的 goroutine 池自动管理 goroutine 的生命周期，
// 可以按需创建，动态缩容。向 goroutine 池提交一个任务，goroutine 池会自动安排某个 goroutine 来处理

type Task struct {
	index int
	nums  []int
	sum   int
	wg    *sync.WaitGroup
}

func (that *Task) Do() {
	for _, num := range that.nums {
		that.sum += num
	}

	that.wg.Done()
}

func taskFunc(data interface{}) {
	task := data.(*Task)
	task.Do()
	fmt.Printf("task:%d sum:%d\n", task.index, task.sum)
}

const (
	DataSize    = 10000
	DataPerTask = 100
)

func main() {
	p, _ := ants.NewPoolWithFunc(10, taskFunc)
	defer p.Release()

	nums := make([]int, DataSize, DataSize)
	for i := range nums {
		nums[i] = rand.Intn(1000)
	}

	var wg sync.WaitGroup
	wg.Add(DataSize / DataPerTask)
	tasks := make([]*Task, 0, DataSize/DataPerTask)
	for i := 0; i < DataSize/DataPerTask; i++ {
		task := &Task{
			index: i + 1,
			nums:  nums[i*DataPerTask : (i+1)*DataPerTask],
			wg:    &wg,
		}
		tasks = append(tasks, task)
		p.Invoke(task)
	}

	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())

	var sum int
	for _, task := range tasks {
		sum += task.sum
	}

	var expect int
	for _, num := range nums {
		expect += num
	}

	fmt.Printf("finish all tasks, result is %d expect:%d\n", sum, expect)
}
