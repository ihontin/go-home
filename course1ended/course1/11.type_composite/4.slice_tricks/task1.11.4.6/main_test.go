package main

import (
	"reflect"
	"testing"
)

type SubSliceCheck struct {
	received []int
	x        []int
	expected []int
}

func tableCheck() []SubSliceCheck {
	var genTable = []SubSliceCheck{
		{[]int{1, 2, 3, 4, 5, 2, 3, 4}, []int{5, 4, 3, 2}, []int{5, 4, 3, 2, 1, 2, 3, 4, 5, 2, 3, 4}},
		{[]int{0, 4, 7, 54, 35}, []int{}, []int{0, 4, 7, 54, 35}},
		{[]int{0, 4, 7, 54, 35}, []int{2, 3}, []int{2, 3, 0, 4, 7, 54, 35}},
		{[]int{}, []int{0, 54}, []int{0, 54}},
	}
	return genTable
}

func TestInsertToStart(t *testing.T) {
	testCheck := tableCheck()
	for _, ch := range testCheck {
		if got := InsertToStart(ch.received, ch.x...); !reflect.DeepEqual(got, ch.expected) {
			t.Errorf("expected = %v, got = %v", ch.expected, got)
		}
	}
}
