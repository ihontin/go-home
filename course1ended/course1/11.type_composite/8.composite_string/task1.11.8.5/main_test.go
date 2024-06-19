package main

import "testing"

type RevCheck struct {
	received, expected string
}

func TestReverseString(t *testing.T) {
	var testTev = []RevCheck{
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"su su", "us us"},
	}
	for _, val := range testTev {
		got := ReverseString(val.received)
		if got != val.expected {
			t.Errorf("expected = %s, got = %s", val.expected, got)
		}
	}
}
