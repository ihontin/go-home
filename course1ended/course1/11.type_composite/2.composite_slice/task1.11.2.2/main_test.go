package main

import "testing"

type SubSliceCheck struct {
	received []int
	expected int
}

func tableCheck() []SubSliceCheck {
	var genTable = []SubSliceCheck{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9},
		{[]int{0, 4, 7, 54, 35, 66}, 66},
		{[]int{0, 0, 7, 8, 7, 8, 0}, 8},
	}
	return genTable
}

func TestMaxDifference(t *testing.T) {
	testCheck := tableCheck()
	for _, check := range testCheck {
		if got := MaxDifference(check.received); got != check.expected {
			t.Errorf("expected = %d, got = %d", check.expected, got)
		}
	}
}
