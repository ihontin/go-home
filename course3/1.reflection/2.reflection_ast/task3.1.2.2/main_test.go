package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()
	_ = w.Close()
	os.Stdout = old

	var byteOut bytes.Buffer
	_, _ = byteOut.ReadFrom(r)

	expected := "func (a *Address) TableName() string { return \"address\" }"
	expected2 := " ID        int     `db:\"id\" db_type:\"INT\" `"
	if !strings.Contains(byteOut.String(), expected) && !strings.Contains(byteOut.String(), expected2) {
		t.Errorf("expected string contains = %s, got = %s", expected, byteOut.String())
	}
}
