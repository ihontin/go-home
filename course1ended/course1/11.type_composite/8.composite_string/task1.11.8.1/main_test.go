package main

import (
	"testing"
)

type CountStr struct {
	received string
	expected int
}

func genTestByte() []CountStr {
	var countByte = []CountStr{
		{"Привет, мир!", 21},
		{"Xnj nfrjt yjdsq ujl", 19},
		{"Что такое новый год", 35},
		{"", 0},
		{"*", 1},
	}
	return countByte
}

func genTestSymbols() []CountStr {
	var countSymbols1 = []CountStr{
		{"Привет, мир!", 12},
		{"Xnj nfrjt yjdsq ujl", 19},
		{"Что такое новый год", 19},
		{"", 0},
		{"*", 1},
	}
	return countSymbols1
}
func TestCountBytes(t *testing.T) {
	var byt = genTestByte()
	for _, b := range byt {
		if got := countBytes(b.received); got != b.expected {
			t.Errorf("expected = %d, got = %d", b.expected, got)
		}
	}
}

func TestCountSymbols(t *testing.T) {
	var sym = genTestSymbols()
	for _, s := range sym {
		if got := countSymbols(s.received); got != s.expected {
			t.Errorf("expected = %d, got = %d", s.expected, got)
		}
	}
}
