package main

import (
	"fmt"
	"time"
)

// Напиши функцию timeout(timeout time.Duration) func() bool, которая будет реализовывать таймауты с помощью каналов.
// Функция timeout является замыканием (closure) и инициализирует канал внутри себя.
func timeout(timeout time.Duration) func() bool {
	ch := make(chan int)
	go func(ch chan int) {
		ch <- 8
	}(ch)
	return func() bool {
		select {
		case <-time.After(timeout):
			return true
		case <-ch:
			return false
		}
	}
}

func main() {
	timeoutFunc := timeout(3 * time.Second)
	since := time.NewTimer(3050 * time.Millisecond)
	for {
		select {
		case <-since.C:
			fmt.Println("Функция не выполнена вовремя")
			return
		default:
			if timeoutFunc() {
				fmt.Println("Функция выполнена вовремя")
				return
			}
		}
	}
}
