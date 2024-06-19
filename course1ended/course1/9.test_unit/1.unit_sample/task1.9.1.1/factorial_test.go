package main

import "testing"

// go test - coverprofile=coverage.out
//go tool cover -func=coverage.out

type FacTests struct {
	n, expected int
}

var testsFact = []FacTests{
	{0, 1},
	{1, 1},
	{5, 120},
}

func TestFactorial(t *testing.T) {
	for _, TestVal := range testsFact {
		if got := Factorial(TestVal.n); got != TestVal.expected {
			t.Errorf("received n = %d, expected = %d, got = %d", TestVal.n, TestVal.expected, got)
		}
	}
}
