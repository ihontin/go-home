package main

import (
	"bytes"
	"os"
	"testing"
)

func testMul(xMul ...interface{}) interface{} {
	var newX int
	if len(xMul) < 1 {
		return newX
	}
	for i, x := range xMul {
		if i == 0 {
			newX = x.(int)
			continue
		}
		newX *= x.(int)
	}
	return newX
}

func testCompare(xMul ...interface{}) interface{} {
	if len(xMul) == 2 {
		return xMul[0] == xMul[1]
	}
	return false
}

type GotExpect struct {
	got      interface{}
	expected interface{}
}

func TestOperate1(t *testing.T) {
	test1 := testCompare
	test2 := testMul
	var testsGot = []GotExpect{
		{funcOperate(test1, 55, 55), true},
		{funcOperate(test2, 5, 5), 25},
		{funcOperate(test1), false},
		{funcOperate(test2, 0, 0), 0},
	}
	for _, testOp := range testsGot {
		if testOp.expected != testOp.got {
			t.Errorf("expected = %v, got = %v", testOp.expected, testOp.got)
		}
	}
}

func TestConcat(t *testing.T) {
	var sumTest = []SumExp{
		{[]interface{}{"2", "3", "5"}, "235"},
		{[]interface{}{}, ""},
		{[]interface{}{"2"}, "2"},
		{[]interface{}{4}, nil},
	}
	for _, st := range sumTest {
		if got := funcConcat(st.received...); got != st.expected {
			t.Errorf("expected = %v, got = %v", st.expected, got)
		}
	}
}

type SumExp struct {
	received []interface{}
	expected interface{}
}

func TestSum(t *testing.T) {
	var sumTest = []SumExp{
		{[]interface{}{2, 3, 5}, 10},
		{[]interface{}{}, nil},
		{[]interface{}{2}, 2},
		{[]interface{}{2.0, 3.0, 5.0}, 10.0},
		{[]interface{}{.0}, 0.0},
		{[]interface{}{true}, nil},
	}
	for _, st := range sumTest {
		if got := funcSum(st.received...); got != st.expected {
			t.Errorf("expected = %v, got = %v", st.expected, got)
		}
	}
}

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	expected := "Hello, World!\n15\n16.5\n"
	var bbuf = bytes.Buffer{}
	bbuf.ReadFrom(r)
	if expected != bbuf.String() {
		t.Errorf("expected = %s, got = %s", expected, bbuf.String())
	}

}
