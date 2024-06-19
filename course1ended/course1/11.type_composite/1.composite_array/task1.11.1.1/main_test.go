package main

import "testing"

var xs = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
var ys = [8]float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5}

func TestSum(t *testing.T) {
	got := sum(xs)
	if expected := 36; expected != got {
		t.Errorf("expected = %d, got = %d", expected, got)
	}

}
func TestAverage(t *testing.T) {
	got := average(xs)
	if expected := 4.5; expected != got {
		t.Errorf("expected = %f, got = %f", expected, got)
	}
}
func TestAverageFloat(t *testing.T) {
	got := averageFloat(ys)
	if expected := 5.0; expected != got {
		t.Errorf("expected = %f, got = %f", expected, got)
	}
}

func TestReverse(t *testing.T) {
	got := reverse(xs)
	if expected := [8]int{8, 7, 6, 5, 4, 3, 2, 1}; expected != got {
		t.Errorf("expected = %v, got = %v", expected, got)
	}
}
