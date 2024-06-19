package main

import "sync"

// Aun для конкурентного безопасного счетчика,
// используя примитивы синхронизации в языке программирования Golang.
type Counter struct {
	count int64
}

// Функция для увеличения значения счетчика на 1
func (c *Counter) Increment() {
	var mutex sync.Mutex
	mutex.Lock()
	c.count += 1
	mutex.Unlock()
}

// Функция для получения текущего значения счетчика
func (c *Counter) GetCount() int64 {
	// Ваш код для получения текущего значения счетчика
	return c.count
}
