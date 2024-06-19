package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

type SubSliceCheck struct {
	received  []int
	expectedX int
	expected  []int
}

func tableCheck() []SubSliceCheck {
	var genTable = []SubSliceCheck{
		{[]int{1, 2, 3, 4, 5, 2, 3, 4}, 1, []int{2, 3, 4, 5, 2, 3, 4}},
		{[]int{0, 4, 7, 54, 35}, 0, []int{4, 7, 54, 35}},
		{[]int{0, 4, 7, 54, 35}, 0, []int{4, 7, 54, 35}},
		{[]int{}, 0, []int{}},
	}
	return genTable
}

func TestPop(t *testing.T) {
	testCheck := tableCheck()
	for _, ch := range testCheck {
		if gotX, got := Pop(ch.received); !reflect.DeepEqual(got, ch.expected) || gotX != ch.expectedX {
			t.Errorf("expected = %v, got = %v, expected Num = %d, got Num = %d,", ch.expected, got, ch.expectedX, gotX)
		}
	}
}

func TestMainFunc(t *testing.T) {
	firstOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = firstOut

	expected := "Значение: 3, Новый срез: [6]"

	var stdout = bytes.Buffer{}
	stdout.ReadFrom(r)
	if expected != stdout.String() {
		t.Errorf("expected = %s, got = %s", expected, stdout.String())
	}

}
