package main

import "testing"

func TestMergeMaps(t *testing.T) {
	map1 := map[string]int{"apple": 3, "banana": 2}
	map2 := map[string]int{"orange": 5, "grape": 4}
	got := mergeMaps(map1, map2)
	expectedValue := 3
	expectedKey := "orange"
	if _, ok := got[expectedKey]; !ok || expectedValue != got["apple"] {
		t.Errorf("expected key = %s, got key= %s, expected value = %d, got value= %d", expectedKey, "", expectedValue, got["apple"])
	}
}
