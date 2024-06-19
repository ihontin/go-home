package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGetReader(t *testing.T) {
	buffer := bytes.NewBufferString("Hello, World!")
	expected := "*bufio.Reader"
	if got := fmt.Sprintf("%v", getReader(buffer)); got != expected {
		fmt.Errorf("expected %s, but got %s", expected, got)
	}
}
