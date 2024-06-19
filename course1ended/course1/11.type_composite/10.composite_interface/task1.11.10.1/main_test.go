package main

import "testing"

func TestGetType(t *testing.T) {
	var genGet = []interface{}{"", 0, []int{}, interface{}(nil)}
	var expected = []string{"string", "int", "[]int", "Пустой интерфейс"}
	for i := 0; i < 4; i++ {
		got := getType(genGet[i])
		if got != expected[i] {
			t.Errorf("expected = %s, got = %s", expected[i], got)
		}
	}
}
