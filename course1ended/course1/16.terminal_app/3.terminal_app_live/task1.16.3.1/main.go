package main

import (
	"fmt"
	"github.com/gosuri/uilive"
	"time"
)

// Текущее время: 20:46:10
// Текущая дата: 2023-08-28
func main() {
	writer := uilive.New()
	writer.Start()
	t := time.Tick(time.Second * 1)
	var ch = make(chan bool)
	go func() {
		time.Sleep(time.Second * 10)
		ch <- true
	}()
	for {
		select {
		case <-t:
			fmt.Fprintf(writer, "Текущее время: %v\n", time.Now().Format("15:04:05"))
			fmt.Fprintf(writer, "Текущая дата: %v\n", time.Now().Format("2006-01-02"))
		case <-ch:
			close(ch)
			writer.Stop()
			return
		}
	}

	writer.Stop() // flush and stop rendering
}
