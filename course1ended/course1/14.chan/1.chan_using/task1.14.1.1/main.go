package main

import (
	"fmt"
	"sync"
	"time"
)

// Эта функция должна принимать переменное количество каналов chans и возвращать канал,
// в который будут отправляться значения из всех каналов chans.
// Функция должна работать до тех пор, пока все каналы chans не будут закрыты.
func mergeChan2(chans ...chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	for _, chain := range chans {
		wg.Add(1)
		go func(chain <-chan int) {
			defer wg.Done()
			for v := range chain {
				out <- v
			}
		}(chain)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Эта функция должна принимать канал mergeTo и переменное количество каналов from.
// Ваша задача - считывать значения из всех каналов from и отправлять их в канал mergeTo.
// Функция должна работать до тех пор, пока все каналы from не будут закрыты.
func mergeChan(mergeTo chan int, from ...chan int) {
	var wg sync.WaitGroup
	wg.Add(len(from))
	for _, fromOne := range from {
		go func(ch chan int) {
			defer wg.Done()
			for val := range ch {
				mergeTo <- val
			}
		}(fromOne)
	}
	wg.Wait()
}
func main() {
	var myChan = make(chan int)
	var chc = make(chan int)
	var chc1 = make(chan int)
	go mergeChan(myChan, chc, chc1)
	go func() {
		defer close(chc)
		for i := 110; i < 115; i++ {
			chc <- i
		}
	}()
	go func() {
		defer close(chc1)
		for i := 0; i < 5; i++ {
			chc1 <- i
		}
	}()
	go func() {
		for val := range myChan {
			fmt.Println(val)
		}
	}()
	time.Sleep(time.Millisecond)

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
	}()

	go func() {
		defer close(ch2)
		ch2 <- 4
		ch2 <- 5
		ch2 <- 6
	}()

	merged := mergeChan2(ch1, ch2)

	for v := range merged {
		fmt.Println(v)
	}

}
