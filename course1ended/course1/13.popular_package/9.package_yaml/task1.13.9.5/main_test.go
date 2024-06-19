package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string `json:"name" yaml:"name"`
	Age  int    `json:"age" yaml:"age"`
}

func TestUnmarshal(t *testing.T) {
	data := []byte(`
people:
  - name: John
    age: 25
  - name: Jane
    age: 30
`)

	expected := struct {
		People []Person `yaml:"people"`
	}{
		People: []Person{
			{Name: "John", Age: 25},
			{Name: "Jane", Age: 30},
		},
	}

	var result struct {
		People []Person `yaml:"people"`
	}

	err := unmarshal(data, &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("unexpected result. got: %v, want: %v", result, expected)
	}
}
