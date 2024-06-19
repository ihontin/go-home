package main

import (
	"fmt"
)

// Необходимо реализовать семафор с помощью каналов в языке программирования Golang. Вам предоставлен код,
// который содержит структуру sema и методы New, Inc и Dec. Ваша задача состоит в том, чтобы реализовать логику семафора,
// чтобы он правильно управлял доступом к ресурсам.
type sema chan struct{}

func New(n int) sema {
	return make(sema, n)
}

func (s sema) Inc(k int) {
	for i := 0; i < k; i++ {
		s <- struct{}{}
	}
}

func (s sema) Dec(k int) {
	for i := 0; i < k; i++ {
		<-s
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	n := len(numbers)

	sem := New(n)

	for _, num := range numbers {
		go func(n int) {
			fmt.Println(n)
			sem.Inc(1)
		}(num)
	}

	sem.Dec(n)
}
