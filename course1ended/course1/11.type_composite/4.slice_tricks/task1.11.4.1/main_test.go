package main

import (
	"reflect"
	"testing"
)

type SubSliceCheck struct {
	received []int
	start    int
	end      int
	expected []int
}

func tableCheck() []SubSliceCheck {
	var genTable = []SubSliceCheck{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 3, []int{2, 3, 4}},
		{[]int{0, 4, 7, 54, 35, 66}, 2, -1, []int{}},
		{[]int{}, 0, 0, []int{}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 15, []int{}},
	}
	return genTable
}

func TestCut(t *testing.T) {
	testCheck := tableCheck()
	for _, ch := range testCheck {
		if got := Cut(ch.received, ch.start, ch.end); !reflect.DeepEqual(got, ch.expected) {
			t.Errorf("expected = %v, got = %v", ch.expected, got)
		}
	}
}
