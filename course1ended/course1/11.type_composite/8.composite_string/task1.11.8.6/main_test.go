package main

import "testing"

type RevCheck struct {
	received string
	expected int
}

func TestCountVowels(t *testing.T) {
	var testTev = []RevCheck{
		{"", 0},
		{"а", 1},
		{"аб", 1},
		{"су су", 2},
	}
	for _, val := range testTev {
		got := CountVowels(val.received)
		if got != val.expected {
			t.Errorf("expected = %d, got = %d", val.expected, got)
		}
	}
}
