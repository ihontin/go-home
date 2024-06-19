package main

import (
	"testing"
	"unicode/utf8"
)

func TestGenerateRandomString(t *testing.T) {
	lenght := []int{10, 0, 6}
	for _, expected := range lenght {
		got := GenerateRandomString(expected)
		if expected != utf8.RuneCountInString(got) {
			t.Errorf("expected = %v, got = %v", expected, len(got))
		}
	}

}
