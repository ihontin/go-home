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

	expected := "&{{<nil>}}\n&{<nil>}\n"
	var bbuf = bytes.Buffer{}
	bbuf.ReadFrom(r)
	if expected != bbuf.String() {
		t.Errorf("expected = %s, got = %s", expected, bbuf.String())
	}

}
