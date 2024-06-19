package main

import "testing"

func TestFibonacciDP(t *testing.T) {
	n := 3
	expected := 2
	if got := FibonacciDP(n); got != expected {
		t.Errorf("expected = %d, got = %d", expected, got)
	}
}
func TestFibonacciBinet(t *testing.T) {
	n := 3
	expected := 2
	if got := FibonacciBinet(n); got != expected {
		t.Errorf("expected = %d, got = %d", expected, got)
	}
}
