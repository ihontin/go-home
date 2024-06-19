package main

import "testing"

func TestFibonacci(t *testing.T) {
	var n = 3
	if got := Fibonacci(n); got != 2 {
		t.Errorf("expected = %d, got = %d", 2, got)
	}
}
