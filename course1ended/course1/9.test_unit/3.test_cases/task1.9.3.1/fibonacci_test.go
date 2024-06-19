package main

import "testing"

type TestsFici struct {
	received, expected int
}

func TestFibonacci(t *testing.T) {
	ficiTest := []TestsFici{
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}
	for _, fi := range ficiTest {
		if got := Fibonacci(fi.received); got != fi.expected {
			t.Errorf("received value = %d expected = %d, got = %d", fi.received, fi.expected, got)
		}
	}
}
