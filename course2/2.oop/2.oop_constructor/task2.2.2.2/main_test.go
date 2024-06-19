package main

import (
	"reflect"
	"testing"
)

func TestWithUsername(t *testing.T) {
	expecteds := []string{"testuser", ""}
	testOrder := &User{}
	for _, expected := range expecteds {
		WithUsername(expected)(testOrder)
		if expected != testOrder.Username {
			t.Errorf("expected = %v, got = %v", expected, testOrder.Username)
		}
	}
}

func TestWithEmail(t *testing.T) {
	expecteds := []string{"testuser@example.com", "item2", ""}
	testOrder := &User{}
	for _, expected := range expecteds {
		WithEmail(expected)(testOrder)
		if !reflect.DeepEqual(expected, testOrder.Email) {
			t.Errorf("expected = %v, got = %v", expected, testOrder.Email)
		}
	}
}

func TestWithRole(t *testing.T) {
	expecteds := []string{"admin", "item2", ""}
	testOrder := &User{}
	for _, expected := range expecteds {
		WithRole(expected)(testOrder)
		if expected != testOrder.Role {
			t.Errorf("expected = %v, got = %v", expected, testOrder.Role)
		}
	}
}

func TestNewUser(t *testing.T) {
	idTest := 21
	nameTest := "testuser"
	EmailTest := "testuser@example.com"
	RoleTest := "admin"
	expected := NewUser(idTest,
		WithUsername(nameTest),
		WithEmail(EmailTest),
		WithRole(RoleTest))
	testO := &User{idTest, nameTest, EmailTest, RoleTest}
	if expected.ID != testO.ID {
		t.Errorf("expected = %v, got = %v", expected.ID, testO.ID)
	}
	if expected.Username != testO.Username {
		t.Errorf("expected = %v, got = %v", expected.Username, testO.Username)
	}
	if expected.Email != testO.Email {
		t.Errorf("expected = %v, got = %v", expected.Email, testO.Email)
	}
	if !reflect.DeepEqual(expected.Role, testO.Role) {
		t.Errorf("expected = %v, got = %v", expected.Role, testO.Role)
	}
}
