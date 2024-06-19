package main

import (
	"reflect"
	"testing"
)

type SubSliceCheck struct {
	received []int
	expCap   int
	expected []int
}

func tableCheck() []SubSliceCheck {
	var genTable = []SubSliceCheck{
		{[]int{1, 2, 3, 4, 5, 2, 3, 4}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{0, 4, 7, 54, 35}, 5, []int{0, 4, 7, 54, 35}},
		{[]int{}, 0, []int{}},
	}
	return genTable
}

func TestRemoveExtraMemory(t *testing.T) {
	testCheck := tableCheck()
	for _, ch := range testCheck {
		if got := RemoveExtraMemory(ch.received); !reflect.DeepEqual(got, ch.expected) || cap(got) != ch.expCap {
			t.Errorf("expected = %v, got = %v, expected cap = %d, got cap = %d", ch.expected, got, ch.expCap, cap(got))
		}
	}
}
