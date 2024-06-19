package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем новый тикер с интервалом 1 секунда
	ticker := time.NewTicker(1 * time.Second)

	data := NotifyEvery(ticker, 5100*time.Millisecond, "Таймер сработал")

	for v := range data {
		fmt.Println(v)
	}

	fmt.Println("Программа завершена")
}

func NotifyEvery(ticker *time.Ticker, d time.Duration, message string) <-chan string {
	tickMess := make(chan string)
	done := make(chan bool)
	go func() {
		time.Sleep(d)
		done <- true
	}()
	go func() {
		for {
			select {
			case <-ticker.C:
				tickMess <- message
			case <-done:
				ticker.Stop()
				close(tickMess)
				return
			}
		}
	}()
	return tickMess
}
