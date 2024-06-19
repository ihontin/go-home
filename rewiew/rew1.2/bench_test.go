package main

import (
	"math/rand"
	"testing"
	"time"
)

func randGener() []int {
	rand.Seed(time.Now().UnixNano())
	numlist := make([]int, 0, 20)
	for i := 0; i < 20; i++ {
		numlist = append(numlist, rand.Intn(20))
	}
	return numlist
}

func BenchmarkFibbRecurs(b *testing.B) {
	benchTests := randGener()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range benchTests {
			FibbRecurs(test)
		}
	}
}
func BenchmarkFibonacciFormula(b *testing.B) {
	benchTests := randGener()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range benchTests {
			FibonacciFormula(test)
		}
	}
}
