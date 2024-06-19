package main

import (
	"os"
	"reflect"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	var newUser = []User{
		{"John Doe", 30, []Comment{
			{"Comment 1"},
			{"Comment 2"},
		},
		},
		{"Jane Smith", 25, []Comment{
			{"Comment 3"},
			{"Comment 4"},
		},
		},
	}
	jsonFilePath, err := os.Getwd()
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	err = writeJSON(jsonFilePath+"/testdata/test.json", newUser)
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	got, err := os.ReadFile(jsonFilePath + "/testdata/test.json")
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	expected := []byte("[{\"name\":\"John Doe\",\"age\":30,\"comments\":[{\"text\":\"Comment 1\"},{\"text\":\"Comment 2\"}]}," +
		"{\"name\":\"Jane Smith\",\"age\":25,\"comments\":[{\"text\":\"Comment 3\"},{\"text\":\"Comment 4\"}]}]")
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected = %v, got = %v", expected, got)
	}
}
