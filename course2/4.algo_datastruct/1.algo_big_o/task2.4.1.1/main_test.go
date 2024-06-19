package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func Test_factorialRecursive(t *testing.T) {
	var a, b int
	for i := 1; i <= 500; i++ {
		a = factorialRecursive(i)
		b = factorialIterative(i)
		if a != b {
			t.Errorf("expected = %d, got = %d", a, b)
		}
	}
}

func Test_factorialIterative(t *testing.T) {
	var a, b int
	for i := 1; i <= 500; i++ {
		a = factorialRecursive(i)
		b = factorialIterative(i)
		if a != b {
			t.Errorf("expected = %d, got = %d", a, b)
		}
	}
}
func Test_compareWhichFactorialIsFaster(t *testing.T) {
	var a, b int
	for i := 1; i <= 500; i++ {
		a = factorialRecursive(i)
		b = factorialIterative(i)
		if a != b {
			t.Errorf("expected = %d, got = %d", a, b)
		}
	}
}

func Test_compareWhichFactorialIsFaster1(t *testing.T) {
	tests := []struct {
		name string
		want map[string]bool
	}{
		{name: "ok",
			want: map[string]bool{"100000": false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareWhichFactorialIsFaster(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compareWhichFactorialIsFaster() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()
	_ = w.Close()
	os.Stdout = old

	expected := "Go version: go1.20.2\nGo OS/Arch: linux / amd64\nWhich factorial is faster?\nmap[100000:false]\n"

	stdOut := bytes.Buffer{}
	_, _ = stdOut.ReadFrom(r)
	if expected != stdOut.String() {
		t.Errorf("expected = %s, got = %s", expected, stdOut.String())
	}
}
