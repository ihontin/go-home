package main

import (
	"bytes"
	"os"
	"testing"
)

var testXs = [8]int{3, 12, 4, 15, 16, 8, 33, 7}
var TestSf = [8]float64{3.1, 22.2, 34.3, 45.4, 56.5, 79.9, 28.8, 37.7}

func TestSortDescInt(t *testing.T) {
	got := sortDescInt(testXs)
	expected := [8]int{33, 16, 15, 12, 8, 7, 4, 3}
	if got != expected {
		t.Errorf("expected = %d, got = %d", expected, got)
	}
}
func TestSortAscInt(t *testing.T) {
	got := sortAscInt(testXs)
	expected := [8]int{3, 4, 7, 8, 12, 15, 16, 33}
	if got != expected {
		t.Errorf("expected = %d, got = %d", expected, got)
	}
}
func TestSortDescFloat(t *testing.T) {
	got := sortDescFloat(TestSf)
	expected := [8]float64{79.9, 56.5, 45.4, 37.7, 34.3, 28.8, 22.2, 3.1}
	if got != expected {
		t.Errorf("expected = %f, got = %f", expected, got)
	}
}
func TestSortAscFloat(t *testing.T) {
	got := sortAscFloat(TestSf)
	expected := [8]float64{3.1, 22.2, 28.8, 34.3, 37.7, 45.4, 56.5, 79.9}
	if got != expected {
		t.Errorf("expected = %f, got = %f", expected, got)
	}
}
func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	var reader = bytes.Buffer{}
	reader.ReadFrom(r)

	expected := "Sorted Int Array (Descending): [9 8 7 5 4 3 2 1]\n" +
		"Sorted Int Array (Ascending): [1 2 3 4 5 7 8 9]\n" +
		"Sorted Float Array (Descending): [9.9 8.8 7.7 5.5 4.4 3.3 2.2 1.1]\n" +
		"Sorted Float Array (Ascending): [1.1 2.2 3.3 4.4 5.5 7.7 8.8 9.9]\n"
	if reader.String() != expected {
		t.Errorf("expected = %s, got = %s", expected, reader.String())
	}

}
