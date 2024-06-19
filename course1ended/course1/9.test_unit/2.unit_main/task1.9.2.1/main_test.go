package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	var reader = bytes.Buffer{}
	reader.ReadFrom(r)

	expected := "Hello, world!\n"
	if reader.String() != expected {
		t.Errorf("expected = %s, got = %s", expected, reader.String())
	}

}
