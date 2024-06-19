package main

import (
	"fmt"
	"sync"
	"time"
)

type someStruct struct {
	name string
}

type someOthestruct struct {
	someStruct
}

func (s *someStruct) Hello() string {
	return s.name
}

/*
1. Goroutines Отличия от потоков
Задачки https://github.com/ansakharov/go_concurrency_tasks

Concurrency is not Parallelism https://www.youtube.com/watch?v=qmg1CF3gZQ0&t=671s
runtime.GOMAXPROCS(1)

https://medium.com/german-gorelkin/race-8936927dba20 серия статей VPN

https://golang-blog.blogspot.com/2019/11/golang-patterns-list.html

2. Channels

count++
make(chan int)
FIFO
1, 2, 3, 4, 5
i, ok := ch

Axioms: https://dave.cheney.net/2014/03/19/channel-axioms

3. Primitives sync / atomic

- var mu sync.Mutex
mu.Lock / mu.Unlock
var mu sync.RWMutex
mu.Rlock / mu.RUnlock

- WaitGroup
var wg sync.WaitGroup
wg.Add() wg.Done() wg.Wait()

sync.Once() Do()
sync.Cond()

sync.Pool() https://habr.com/ru/articles/277137/

4. Context
https://habr.com/ru/companies/nixys/articles/461723/

ctx := context.Background() / context.TODO()
ctx, cancel := context.WithTimeout(ctx, 10)
*/

func main() {
	counter := 20
	for i := 0; i < counter; i++ {
		go func(i int) {
			fmt.Println(i * i)
		}(i)
	}

	time.Sleep(time.Second)
}

func merge(chans ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	result := make(chan int)

	for _, singleCh := range chans {
		wg.Add(1)
		singleCh := singleCh
		go func() {
			defer wg.Done()
			for val := range singleCh {
				result <- val
			}
		}()
	}
	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}
