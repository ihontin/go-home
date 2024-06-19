package main

import "testing"

type expGot struct {
	receive interface{}
	expect  string
}

func TestGetVariableType(t *testing.T) {
	genTest := []expGot{
		{true, "bool"},
		{2.0, "float64"},
		{2, "int"},
		{map[string]interface{}{"a": 1}, "map[string]interface {}"},
	}
	for _, tog := range genTest {
		if got := getVariableType(tog.receive); got != tog.expect {
			t.Errorf("expected = %s, got = %s", tog.expect, got)
		}
	}
}
