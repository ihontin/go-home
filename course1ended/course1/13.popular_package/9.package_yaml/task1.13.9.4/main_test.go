package main

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

type Person struct {
	Name string `json:"name" yaml:"name"`
	Age  int    `json:"age" yaml:"age"`
}

func TestWriteYAML(t *testing.T) {
	filePath := "test.yaml"
	expected := struct {
		People []Person `yaml:"people"`
	}{
		People: []Person{
			{Name: "John", Age: 25},
			{Name: "Jane", Age: 30},
		},
	}

	err := writeYAML(filePath, &expected)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result struct {
		People []Person `yaml:"people"`
	}

	err = yaml.Unmarshal(fileData, &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("unexpected result. got: %v, want: %v", result, expected)
	}

	err = os.Remove(filePath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
