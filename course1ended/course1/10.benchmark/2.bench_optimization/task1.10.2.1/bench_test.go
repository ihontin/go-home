package main

import "testing"

func BenchmarkFibonacciDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciDP(3)
	}
}
func BenchmarkFibonacciBinet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciDP(3)
	}
}
