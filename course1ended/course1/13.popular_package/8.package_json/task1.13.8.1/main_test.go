package main

import (
	"fmt"
	"testing"
)

func TestGetJSON(t *testing.T) {
	var newUser = []User{
		{"Alex", 46, []Comment{
			{"what can i say"},
			{"everithing"},
		},
		},
		{"Susanna", 14, []Comment{
			{"Shiny star"},
			{"a cup of milk"},
		},
		},
	}
	got, err := getJSON(newUser)
	if err != nil {
		fmt.Errorf("%v", err.Error())
	}
	expected := "[{\"name\":\"Alex\",\"age\":46,\"comments\":[{\"text\":\"what can i say\"},{\"text\":\"everithing\"}]}," +
		"{\"name\":\"Susanna\",\"age\":14,\"comments\":[{\"text\":\"Shiny star\"},{\"text\":\"a cup of milk\"}]}]"
	if expected != got {
		fmt.Printf("expected = %s, got = %s", expected, got)
	}
	got2, err := getJSON([]User{})
	if err != nil {
		fmt.Errorf("%v", err.Error())
	}
	expected2 := "[]"
	if expected != got {
		fmt.Printf("expected = %s, got = %s", expected2, got2)
	}
}
