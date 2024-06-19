package main

import (
	"reflect"
	"testing"
)

type SubSliceCheck struct {
	received []int
	ind      int
	insert   []int
	expected []int
}

func tableCheck() []SubSliceCheck {
	var genTable = []SubSliceCheck{
		{[]int{1, 2, 3, 4, 5}, 2, []int{6, 7, 8}, []int{1, 2, 3, 6, 7, 8, 4, 5}},
		{[]int{0, 4, 7, 54, 35, 66}, -2, []int{}, []int{}},
		{[]int{}, 0, []int{}, []int{}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 15, []int{6, 7, 8}, []int{}},
	}
	return genTable
}

func TestInsertAfterIDX(t *testing.T) {
	testCheck := tableCheck()
	for _, ch := range testCheck {
		if got := InsertAfterIDX(ch.received, ch.ind, ch.insert...); !reflect.DeepEqual(got, ch.expected) {
			t.Errorf("expected = %d, got = %d", ch.expected, got)
		}
	}
}
