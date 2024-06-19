package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetUsersFromJSON(t *testing.T) {
	jsonData := []byte(`[
		{
			"name": "John",
			"age": 30,
			"comments": [
				{"text": "Great post!"},
				{"text": "I agree"}
			]
		},
		{
			"name": "Alice",
			"age": 25,
			"comments": [
				{"text": "Nice article"}
			]
		}
	]`)
	users, err := getUsersFromJSON(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var expected = []interface{}{"John", 30, "Great post!", "I agree", "Alice", 25, "Nice article"}
	var gotNameAge []interface{}
	for _, user := range users {
		gotNameAge = append(gotNameAge, user.Name)
		gotNameAge = append(gotNameAge, user.Age)
		for _, comment := range user.Comments {
			gotNameAge = append(gotNameAge, comment.Text)
		}
	}
	if !reflect.DeepEqual(gotNameAge, expected) {
		t.Errorf("expected = %s,\n got = %s", expected, gotNameAge)
	}
}
