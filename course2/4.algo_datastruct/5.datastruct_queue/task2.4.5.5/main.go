package main

import "fmt"

type CircuitRinger interface {
	Add(val int)
	Get() (int, bool)
}
type RingBuffer struct {
	queue []int
	size  int
}

func NewRingBuffer(n int) *RingBuffer {
	return &RingBuffer{queue: make([]int, 0, n), size: n}
}

func (b *RingBuffer) Add(val int) {
	if len(b.queue) < b.size {
		b.queue = append(b.queue, val)
	} else {
		b.queue = append(b.queue[1:], val)
	}

}

func (b *RingBuffer) Get() (int, bool) {
	if len(b.queue) == 0 {
		return 0, false
	}
	getVal := b.queue[0]
	b.queue = b.queue[1:]
	return getVal, true
}

func main() {
	rb := NewRingBuffer(3)
	rb.Add(1)
	rb.Add(2)
	rb.Add(3)
	rb.Add(4)

	if val, ok := rb.Get(); ok {
		fmt.Println(val) // Выводит: 2
	}
	if val, ok := rb.Get(); ok {
		fmt.Println(val) // Выводит: 3
	}
	if val, ok := rb.Get(); ok {
		fmt.Println(val) // Выводит: 4
	}
	if _, ok := rb.Get(); !ok {
		fmt.Println("Буфер пуст") // Выводит: Буфер пуст
	}
}
