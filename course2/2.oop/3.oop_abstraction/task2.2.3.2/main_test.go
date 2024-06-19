package main

import (
	"testing"
)

func TestSQLiteGenerator_CreateTableSQL(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("CreateTableSQL not panicked: %v", r)
		}
	}()
	sqTable := &SQLiteGenerator{}
	var tableN Tabler

	got := sqTable.CreateTableSQL(tableN)
	expected := "implement me"
	if got != expected {
		t.Errorf("expected: %vб got: %v", expected, got)
	}
}
func TestSQLiteGenerator_CreateInsertSQL(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("CreateTableSQL not panicked: %v", r)
		}
	}()
	sqTable := &SQLiteGenerator{}
	var tableN Tabler

	got := sqTable.CreateInsertSQL(tableN)
	expected := "implement me"
	if got != expected {
		t.Errorf("expected: %vб got: %v", expected, got)
	}
}
func TestGoFakeitGenerator_TableName(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("CreateTableSQL not panicked: %v", r)
		}
	}()
	sqTable := &GoFakeitGenerator{}

	got := sqTable.TableName()
	expected := "implement me"
	if got != expected {
		t.Errorf("expected: %vб got: %v", expected, got)
	}
}
func TestGoFakeitGenerator_GenerateFakeUser(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("CreateTableSQL not panicked: %v", r)
		}
	}()
	sqTable := &GoFakeitGenerator{}

	got := sqTable.GenerateFakeUser()
	expected := User{
		0, "", "", "",
	}
	if got.ID != expected.ID {
		t.Errorf("expected: %vб got: %v", expected, got)
	}
}
