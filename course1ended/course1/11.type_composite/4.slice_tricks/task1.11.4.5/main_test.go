package main

import (
	"reflect"
	"testing"
)

type SubSliceCheck struct {
	received []int
	divider  int
	expected []int
}

func tableCheck() []SubSliceCheck {
	var genTable = []SubSliceCheck{
		{[]int{1, 2, 3, 4, 5, 2, 3, 4}, 5, []int{5}},
		{[]int{0, 4, 7, 54, 35}, 0, []int{0, 4, 7, 54, 35}},
		{[]int{0, 4, 7, 54, 35}, 6, []int{0, 54}},
		{[]int{}, 0, []int{}},
	}
	return genTable
}

func TestFilterDividers(t *testing.T) {
	testCheck := tableCheck()
	for _, ch := range testCheck {
		if got := FilterDividers(ch.received, ch.divider); !reflect.DeepEqual(got, ch.expected) {
			t.Errorf("expected = %v, got = %v", ch.expected, got)
		}
	}
}
