package main

import (
	"reflect"
	"testing"
)

type SubSliceCheck struct {
	received []int
	expected []int
}

func tableCheck() []SubSliceCheck {
	var genTable = []SubSliceCheck{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{0, 0, 4, 1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{0, 4, 7, 54, 35, 66}, []int{0, 0, 4, 1, 0, 0, 4, 7, 54, 35, 66}},
		{[]int{0, 0, 7, 8, 7, 8, 0}, []int{0, 0, 4, 1, 0, 0, 0, 7, 8, 7, 8, 0}},
	}
	return genTable
}

func TestAppendInt(t *testing.T) {
	testCheck := tableCheck()
	var baseList = []int{0, 0, 4, 1, 0}
	for _, check := range testCheck {
		if got := appendInt(baseList, check.received...); !reflect.DeepEqual(got, check.expected) {
			t.Errorf("expected = %d, got = %d", check.expected, got)
		}
	}
}
