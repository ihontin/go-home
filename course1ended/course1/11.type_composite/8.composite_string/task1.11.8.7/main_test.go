package main

import (
	"bytes"
	"os"
	"testing"
)

func TestReplaceSymbols(t *testing.T) {
	var testTev, old, newO = "Hello, world!", 'o', '0'
	got := ReplaceSymbols(testTev, old, newO)
	expected := "Hell0, w0rld!"
	if got != expected {
		t.Errorf("expected = %s, got = %s", expected, got)
	}
}

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	expected := "Hell0, w0rld!\n"
	var stdout = bytes.Buffer{}
	stdout.ReadFrom(r)
	if expected != stdout.String() {
		t.Errorf("expected = %s, got = %s", expected, stdout.String())
	}

}
