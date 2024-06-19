package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {

	oldStdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdin = r
	os.Stdout = w

	input := "input data"
	go func() {
		defer w.Close()
		fmt.Fprint(w, input)
	}()

	main()

	w.Close()

	os.Stdout = oldStdout

	var stdout bytes.Buffer
	stdout.ReadFrom(r)

	expected := "expected output"
	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
