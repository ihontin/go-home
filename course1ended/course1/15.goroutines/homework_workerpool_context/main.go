package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	workerPool()
}

func workerPool() {
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*20)
	defer cancel()

	var wg = &sync.WaitGroup{}
	ch1, ch2 := make(chan int), make(chan int)
	for i := 0; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, ch1, ch2)
		}()
	}
	go func() {
		for i := 0; i < 1000; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		wg.Wait()
		close(ch2)
	}()
	var count int
	for val := range ch2 {
		fmt.Println(val)
		count++
	}
	fmt.Println(count)
}

func worker(ctx context.Context, ch1 <-chan int, ch2 chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-ch1:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			ch2 <- val
		}
	}
}
