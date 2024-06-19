package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"testing"
)

func TestGetUniqueUsers(t *testing.T) {
	var users []User
	for i := 0; i < 1000; i++ {
		users = append(users, User{gofakeit.FirstName(), gofakeit.Number(10, 120), gofakeit.Email()})
	}
	var numUsers = make(map[string]int)
	for _, us := range users {
		numUsers[us.Nickname]++
	}
	expected := len(numUsers)
	got := cap(getUniqueUsers(users))
	if expected != got {
		t.Errorf("cap() should be equil len(), but: expected len = %d, got cap = %d", expected, got)
	}
}
