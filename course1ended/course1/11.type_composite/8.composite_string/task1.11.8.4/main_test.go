package main

import "testing"

func TestConcatStrings(t *testing.T) {
	a := []string{"All", " ", "next", " ", "month", "!"}
	var b []string
	c := []string{"", "", "", "", "", ""}
	var got, expected string
	for i := 0; i < 3; i++ {
		switch i {
		case 0:
			got = concatStrings(a...)
			expected = "All next month!"
		case 1:
			got = concatStrings(b...)
			expected = ""
		case 2:
			got = concatStrings(c...)
			expected = ""
		}
		if got != expected {
			t.Errorf("expected = %s, got = %s", expected, got)
		}
	}
}
