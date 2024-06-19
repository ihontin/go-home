package main

import (
	"bytes"
	"os"
	"testing"
)

type SubSliceCheck struct {
	received, expected []int
}

func tableCheck() []SubSliceCheck {
	var genTable = []SubSliceCheck{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{3, 4, 5, 6}},
		{[]int{0, 4, 7, 54, 35, 66}, []int{7, 54, 35, 66}},
		{[]int{0, 0, 7, 8, 7, 8, 0}, []int{7, 8, 7, 8}},
	}
	return genTable
}

func TestGetSubSlice(t *testing.T) {
	listVal := tableCheck()
	for _, check := range listVal {
		got := getSubSlice(check.received, 2, 6)
		for i, _ := range got {
			if got[i] != check.expected[i] {
				t.Errorf("received = %v, expected = %v", listVal, got)
			}
		}
	}
}

func TestMainFunc(t *testing.T) {
	old := os.Stdout     // в переменную записыаешь поток вывода
	r, w, _ := os.Pipe() //открываешь поток жестко
	os.Stdout = w        // вывод жестко может записывать

	main() // вывзо мэйна

	w.Close()       // закрываешь поток записи
	os.Stdout = old // в вывод записываешь чё там было

	expected := "[3 4 5 6]\n"
	var stdOut bytes.Buffer
	stdOut.ReadFrom(r)
	if stdOut.String() != expected {
		t.Errorf("expected = %v, got = %v", expected, stdOut.String())
	}

}
