package main

import (
	"fmt"
	"sync"
	"time"
)

// Необходимо реализовать функцию waitGroupExample, которая будет принимать в качестве аргументов функции goroutines типа
// func(n int) string. Функции goroutines должны выполняться параллельно. Каждая функция goroutines должна возвращать
// строку в формате “goroutine N done”, где N - номер горутины.
//
// Функция waitGroupExample должна возвращать результаты выполнения горутин, объединенные в одну строку,
// где каждый результат находится на отдельной строке.
func waitGroupExample(goroutines ...func() string) string {
	var result = make(chan string)
	var allInOne string
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i, f := range goroutines {
		wg.Add(1)
		go func(f func() string, i int) {
			defer wg.Done()
			if i == 0 {
				time.Sleep(time.Millisecond * 100)
			}
			mutex.Lock()
			result <- f() + "\n"
			mutex.Unlock()
		}(f, i)
	}
	for {
		select {
		case res := <-result:
			allInOne += res
		case <-time.After(time.Second * 1):
			wg.Wait()
			close(result)
			return allInOne
		}
	}
}
func main() {
	count := 1000
	goroutines := make([]func() string, count)
	for i := 0; i < count; i++ {
		j := i
		goroutines[j] = func() string {
			return fmt.Sprintf("goroutine %d", j)
		}
	}
	fmt.Println(waitGroupExample(goroutines...))
}
