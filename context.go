package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func play(ctx context.Context) <-chan int {
	dist := make(chan int)
	n := 1

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dist <- n:
				n++
			}
		}
	}()

	return dist
}

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting...")
		time.Sleep(time.Millisecond * 10) // 假设正常逻辑为10ms
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done")
	wg.Done()
}

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//
	//for n := range play(ctx) {
	//	fmt.Println(n)
	//	if n == 5 {
	//		break
	//	}
	//}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
