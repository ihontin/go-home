package main

import "testing"

type ValidCount struct {
	received string
	expected int
}

func TestCountUniqueUTF8Chars(t *testing.T) {
	var checkE = []ValidCount{
		{"invalid_email", 9},
		{"0", 1},
		{"aaa****)))", 3},
		{"", 0},
	}
	for _, exp := range checkE {
		if got := countUniqueUTF8Chars(exp.received); got != exp.expected {
			t.Errorf("expected = %d, got = %d", exp.expected, got)
		}
	}

}
