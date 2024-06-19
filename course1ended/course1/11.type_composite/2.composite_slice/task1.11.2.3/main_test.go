package main

import (
	"bytes"
	"os"
	"testing"
)

type SubSliceCheck struct {
	received []int
	expected int
}

func tableCheck() []SubSliceCheck {
	var genTable = []SubSliceCheck{
		{[]int{1, 2, 3, 4, 9, 4, 3, 2, 1}, 9},
		{[]int{1, 2, 3, 4, 66, 4, 3, 2, 1, 10}, 66},
		{[]int{9, 1, 2, 3, 4, 5, 4, 3, 2, 1}, 9},
	}
	return genTable
}

func TestMaxDifference(t *testing.T) {
	testCheck := tableCheck()
	for _, check := range testCheck {
		if got := findSingleNumber(check.received); got != check.expected {
			t.Errorf("expected = %d, got = %d", check.expected, got)
		}
	}
}

func TestBitwiseXOR(t *testing.T) {
	type XorCheck struct {
		a, b, expected int
	}
	var testBit = []XorCheck{
		{3, 3, 0},
		{5, 10, 15},
		{3, 4, 7},
		{3, 5, 6},
		{5, 1, 4},
	}
	for _, check := range testBit {
		if got := bitwiseXOR(check.a, check.b); got != check.expected {
			t.Errorf("expected = %d, got = %d", check.expected, got)
		}
	}
}

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	var stdout = bytes.Buffer{}
	stdout.ReadFrom(r)
	expected := "5\n"
	if expected != stdout.String() {
		t.Errorf("expected = %s, got = %s", expected, stdout.String())
	}
}
