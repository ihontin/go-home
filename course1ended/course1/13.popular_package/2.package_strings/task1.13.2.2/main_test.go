package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestMainFanc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	expected := "[{Betty [{good Comment 1} {Use camelCase please}]} {Jhon [{Good Comment 1} {Good Comment 2} {Good Comment 3}]}]\n"
	var outStd bytes.Buffer
	outStd.ReadFrom(r)
	if expected != outStd.String() {
		t.Errorf("expected = %s, got = %s", expected, outStd.String())
	}
}

func genUsers() []User {
	users := []User{
		{
			Name: "Betty",
			Comments: []Comment{
				{Message: "good Comment 1"},
				{Message: "BaD CoMmEnT 2"},
				{Message: "Bad Comment 3"},
				{Message: "Use camelCase please 4"},
				{Message: "Bad Comment Good Comment 5"},
			},
		},
		{
			Name: "Jhon",
			Comments: []Comment{
				{Message: "Good Comment 1"},
				{Message: "Good Comment 2"},
				{Message: "Good Comment 3"},
				{Message: "Bad Comments 4"},
				{Message: ""},
			},
		},
	}
	return users
}
func TestFilterComments(t *testing.T) {
	users := genUsers()
	got := FilterComments(users)
	var expected = [][]string{{"good Comment 1", "Use camelCase please 4"}, {"Good Comment 1", "Good Comment 2", "Good Comment 3", ""}}
	for i, user := range got {
		for j, coment := range user.Comments {
			if coment.Message != expected[i][j] {
				t.Errorf("expected = %s, got = %s", expected[i][j], coment.Message)
			}
		}
	}
}

func TestIsBadComment(t *testing.T) {
	var coments = [][]bool{{false, true, true, false, true}, {false, false, false, true, false}}
	users := genUsers()
	for i, user := range users {
		for j, coment := range user.Comments {
			if got := IsBadComment(coment.Message); got != coments[i][j] {
				t.Errorf("expected = %t, got = %t", coments[i][j], got)
			}
		}
	}
}

func TestGetBadComments(t *testing.T) {
	var expected = []Comment{
		{"BaD CoMmEnT 2"},
		{"Bad Comment 3"},
		{"Bad Comment Good Comment 5"},
		{"Bad Comments 4"},
	}
	users := genUsers()
	got := GetBadComments(users)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected = %v, got = %v", expected, got)
	}
}

func TestGetGoodComments(t *testing.T) {
	var expected = []Comment{
		{"good Comment 1"},
		{"Bad Comment Good Comment 5"},
		{"Good Comment 1"},
		{"Good Comment 2"},
		{"Good Comment 3"},
	}
	users := genUsers()
	got := GetGoodComments(users)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected = %v, got = %v", expected, got)
	}
}
