package main

import "testing"

type expGot struct {
	receive1 []int
	receive2 string
	expect   string
}

func TestGenerateMathString(t *testing.T) {
	genTest := []expGot{
		{[]int{}, "+", ""},
		{[]int{3, 97}, "", ""},
		{[]int{}, "", ""},
		{[]int{3, 97}, "+", "3 + 97 = 100"},
	}
	for _, tog := range genTest {
		if got := generateMathString(tog.receive1, tog.receive2); got != tog.expect {
			t.Errorf("expected = %s, got = %s", tog.expect, got)
		}
	}

}
