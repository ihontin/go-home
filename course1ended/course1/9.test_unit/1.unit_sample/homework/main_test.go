package main

import (
	"testing"
)

// go test - coverprofile=coverage.out
// go tool cover -func=coverage.out
type AddTest struct {
	a, b, expected int
}

var addTests = []AddTest{
	{3, 3, 6},
	{8, 2, 10},
	{60, 43, 103},
	{21, 21, 42},
	{56, 44, 100},
}

func TestAdd(t *testing.T) {
	for _, values := range addTests {
		if got := Add(values.a, values.b); got != values.expected {
			t.Errorf("received a = %d, b = %d, expected = %d, got = %d", values.a, values.b, values.expected, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	var (
		a = 28
		b = 12
	)
	expected := 16
	got := Subtract(a, b)
	if expected != got {
		t.Errorf("received a = %d, b = %d, expected = %d, got = %d", a, b, expected, got)
	}

}
