package main

import (
	"testing"
)

func TestUser_TableName(t *testing.T) {
	var userT = []User{
		{13, "Loca", "Pava", "Pava@i.ru"},
		{0, "", "", ""},
		{2, "1", "2", "3@i.ru"},
	}
	for _, val := range userT {
		got := val.TableName()
		if got.ID != val.ID || got.FirstName != val.FirstName || got.Email != val.Email {
			t.Errorf("expected = %v, got = %v\n", val, got)
		}
	}
}

func TestSQLiteGenerator_CreateTableSQL(t *testing.T) {
	var userT = []User{
		{13, "Loca", "Pava", "Pava@i.ru"},
		{0, "", "", ""},
	}
	var expected = []string{
		"CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, first_name VARCHAR(100), last_name VARCHAR(100), email VARCHAR(100) UNIQUE);",
		"CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, first_name VARCHAR(100), last_name VARCHAR(100), email VARCHAR(100) UNIQUE);",
	}
	sqlGenerator := &SQLiteGenerator{}
	for i, val := range userT {
		got := sqlGenerator.CreateTableSQL(&val)
		if got != expected[i] {
			t.Errorf("expected = %s, got = %s", expected, got)
		}
	}
}

func TestSQLiteGenerator_CreateInsertSQL(t *testing.T) {
	var userT = []User{
		{13, "Loca", "Pava", "Pava@i.ru"},
		{0, "", "", ""},
		{2, "1", "2", "3@i.ru"},
	}
	var expected = []string{
		"INSERT INTO users (id, first_name, last_name, email) VALUES (13,'Loca','Pava','Pava@i.ru');",
		"INSERT INTO users (id, first_name, last_name, email) VALUES (0,'','','');",
		"INSERT INTO users (id, first_name, last_name, email) VALUES (2,'1','2','3@i.ru');",
	}
	sqlGenerator := &SQLiteGenerator{}
	for i, val := range userT {
		got := sqlGenerator.CreateInsertSQL(&val)
		if got != expected[i] {
			t.Errorf("expected = %s, got = %s", expected, got)
		}
	}
}

func TestGoFakeitGenerator_GenerateFakeUser(t *testing.T) {
	sqlGenerator := &SQLiteGenerator{}
	for i := 0; i < 3; i++ {
		got := sqlGenerator.gfg.GenerateFakeUser()
		expected := got.TableName()
		if expected != &got {
			t.Errorf("expected = %v, got = %v", expected, got)
		}
	}
}

func TestGenerateUserInserts(t *testing.T) {
	gotList := GenerateUserInserts(3)
	expected := "INSERT INTO users (id, first_name, last_name, email) VALUES (0,'','','');"
	for _, got := range gotList {
		if got == "" {
			t.Errorf("expected = %v, got = %v", expected, got)
		}
	}
}
