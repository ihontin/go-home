package main

import (
	"fmt"
	"testing"
)

type ValidCount struct {
	received string
	expected string
}

func TestIsRussianLetter(t *testing.T) {
	var checkLett = map[rune]bool{
		0:    false,
		1041: true,
		1026: false,
	}
	for k, v := range checkLett {
		if got := isRussianLetter(k); got != v {
			t.Errorf("expected = %v, got = %v", v, got)
		}
	}

}

func TestCountRussianLetters(t *testing.T) {

	var checkE = []ValidCount{
		{"Привет, мир!", "map[1074:1 1077:1 1080:2 1084:1 1087:1 1088:2 1090:1]"},
		{"!****)))", "map[]"},
		{"", "map[]"},
	}
	for _, exp := range checkE {
		if got := fmt.Sprintf("%v", countRussianLetters(exp.received)); got != exp.expected {
			t.Errorf("expected = %v, got = %v", exp.expected, got)
		}
	}
}
