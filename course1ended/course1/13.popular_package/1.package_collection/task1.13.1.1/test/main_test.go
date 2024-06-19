package main

import (
	"testing"
)

func genUsers() []*User {
	users := []*User{
		{Id: 1, Name: "John", Age: 20, Nickname: "Purple"},
		{Id: 2, Name: "Tom", Age: 22, Nickname: "Green"},
		{Id: 3, Name: "Billy", Age: 20, Nickname: "Black"},
		{Id: 4, Name: "Mister X", Age: 30, Nickname: "Green"},
	}
	return users
}

type ByCond struct {
	received string
	expected []string
}

func TestGetUsersByCondition(t *testing.T) {
	var allUs = genUsers()
	var byCondi = []ByCond{
		{"age > 18", []string{"John", "Tom", "Billy", "Mister X"}},
		{"age < 18", []string{}},
		{"age >= 18", []string{"John", "Tom", "Billy", "Mister X"}},
		{"age <= 18", []string{}},
		{"age = 18", []string{}},
		{"", nil},
	}
	for _, us := range byCondi {
		getByCondition := getUsersByCondition(allUs, us.received)
		for i, n := range getByCondition {
			if (*n).Name != us.expected[i] {
				t.Errorf("expected = %s, got = %s", us.expected[i], (*n).Name)
			}
		}
	}
}

type ByAg struct {
	received int
	expected []string
}

func TestGetUsersByAge(t *testing.T) {
	var allUs = genUsers()
	var byCondi = []ByAg{
		{30, []string{"Mister X"}},
		{22, []string{"Tom"}},
		{20, []string{"John", "Billy"}},
		{0, []string{}},
	}
	for _, us := range byCondi {
		getByCondition := getUsersByAge(allUs, us.received)
		for i, n := range getByCondition {
			if (*n).Name != us.expected[i] {
				t.Errorf("expected = %s, got = %s", us.expected[i], (*n).Name)
			}
		}
	}
}

type ByNic struct {
	received string
	expected []string
}

func TestGetUsersByNickName(t *testing.T) {
	var allUs = genUsers()
	var byCondi = []ByNic{
		{"Green", []string{"Tom", "Mister X"}},
		{"Tom", []string{}},
		{"", []string{}},
	}
	for _, us := range byCondi {
		getByCondition := getUsersByNickName(allUs, us.received)
		for i, n := range getByCondition {
			if (*n).Name != us.expected[i] {
				t.Errorf("expected = %s, got = %s", us.expected[i], (*n).Name)
			}
		}
	}
}

func TestGetUsersUniqueNickName(t *testing.T) {
	var allUs = genUsers()
	getByCondition := getUsersUniqueNickName(allUs)
	var expected = []string{"John", "Tom", "Billy"}
	for i, n := range getByCondition {
		if (*n).Name != expected[i] {
			t.Errorf("expected = %s, got = %s", expected[i], (*n).Name)
		}
	}
}
