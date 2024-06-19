package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGetDataString(t *testing.T) {
	buffer := bytes.NewBufferString("Hello, World!")
	expected := "Hello, World!"
	if got := fmt.Sprintf("%v", getDataString(buffer)); got != expected {
		fmt.Errorf("expected %s, but got %s", expected, got)
	}
}
