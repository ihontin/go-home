package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 1. Задача  Написать функцию, которая принимает канал и число N типа int,
//после получения N значений из канала, функция должна вернуть срез с этими значениями

func rFromChan(ch <-chan int, n int) []int {
	var s = make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, <-ch)
	}
	return s
}

func RandNumbers(length, max int) []int {
	var s = make([]int, 0, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		s = append(s, rand.Intn(max))
	}
	return s
}

func writeToChan(ch chan<- int) {
	defer close(ch)
	for _, v := range RandNumbers(100, 100) {
		ch <- v
	}
}

func mergeChan(ch ...chan int) chan int {
	outch := make(chan int)
	var wg sync.WaitGroup
	for _, singleCh := range ch {
		wg.Add(1)
		singleCh := singleCh
		go func() {
			defer wg.Done()
			for val := range singleCh {
				outch <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outch)
	}()

	return outch
}

func main() {
	nChan := make(chan int)
	n := 15
	go func() {
		for i := 0; i < n; i++ {
			nChan <- i
		}
	}()
	s := rFromChan(nChan, 10)

	fmt.Println(s)

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	mergedChan := mergeChan(ch1, ch2, ch3, ch4)
	go writeToChan(ch1)
	go writeToChan(ch2)
	go writeToChan(ch3)
	go writeToChan(ch4)

	for val := range mergedChan {
		fmt.Println(val)
	}
}
