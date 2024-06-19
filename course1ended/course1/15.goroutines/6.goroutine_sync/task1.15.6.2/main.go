package main

import (
	"sync"
)

type Counter struct {
	value int
}

func (c *Counter) Increment() int {
	c.value += 1
	return c.value
}

// concurrentSafeCounter, которая будет представлять собой конкурентно-безопасный
// счетчик с использованием sync.Mutex. Функция должна иметь следующую сигнатуру:
func concurrentSafeCounter() int {
	counter := Counter{0}
	var wg sync.WaitGroup
	var mut sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mut.Lock()
			counter.Increment()
			mut.Unlock()
		}()
	}
	wg.Wait()
	return counter.value
}

//func main() {
//	fmt.Println(concurrentSafeCounter())
//	fmt.Println(concurrentSafeCounter())
//}
