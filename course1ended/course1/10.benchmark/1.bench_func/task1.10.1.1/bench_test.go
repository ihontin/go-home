package main

import "testing"

// go test -bench=. -benchmem -benchtime=1000x
//аллокации сбор статистики распределения памяти
// b.ReportAllocs()

//b.ResetTimer()

//Профилирование CPU
//go test -bench=. -cpuprofile=cpu.prof
//go tool pprof cpu.prof

//Профилирование памяти
//go test -bench=. -memprofile=mem.prof
//go tool pprof --alloc_space mem.prof

//Профилирование блокировок: Профилирование блокировок помогает идентифицировать горутины,
//которые проводят значительное время в состоянии блокировки (ожидания).
//go test -bench=. -blockprofile=block.prof
//go tool pprof --block block.prof

func BenchmarkMyFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			Fibonacci(j)
		}
	}
}
