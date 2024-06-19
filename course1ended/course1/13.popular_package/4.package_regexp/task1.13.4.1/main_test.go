package main

import "testing"

type ValidEmail struct {
	received string
	expected bool
}

func TestIsValidEmail(t *testing.T) {
	var checkE = []ValidEmail{
		{"invalid_email", false},
		{"test@example.com", true},
		{"test`example.com", true},
		{"", false},
	}
	for _, exp := range checkE {
		if got := isValidEmail(exp.received); got != exp.expected {
			t.Errorf("expected = %t, got = %t", exp.expected, got)
		}
	}
}
