package main

import (
	"database/sql"
	"log"
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

func TestNewMigrator(t *testing.T) {
	db, err := sql.Open("sqlite3", "file:my_database.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Создание мигратора с использованием вашего SQLGenerator
	YourSQLGeneratorInstance := &SQLiteGenerator{}
	got := NewMigrator(db, YourSQLGeneratorInstance)
	if got.db != db {
		t.Errorf("expected = %v, got = %v", db, got.db)
	}
	if got.sqlGenerator != YourSQLGeneratorInstance {
		t.Errorf("expected = %v, got = %v", YourSQLGeneratorInstance, got.sqlGenerator)
	}
}

func TestMigrator_Migrate(t *testing.T) {
	db, err := sql.Open("sqlite3", "file:my_database.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)

	}
	defer db.Close()

	// Создание мигратора с использованием вашего SQLGenerator
	YourSQLGeneratorInstance := &SQLiteGenerator{}
	migrator := NewMigrator(db, YourSQLGeneratorInstance)

	// Миграция таблицы User
	newUser := &User{5, "tou", "you", "boy"}
	err = migrator.Migrate(newUser)
	if err != nil {
		t.Errorf("error not expected = %v", err)
		//log.Fatalf("failed to migrate: %v", err)
	}

	rows, err := db.Query("PRAGMA table_info(users)")
	if err != nil {
		log.Fatalf("QueryRow error: %v", err)
	}
	defer rows.Close()
	var indexCount int
	expected := []string{"id", "first_name", "last_name", "email"}
	for rows.Next() {
		var count int
		var got, LastName, fieldNil string
		var dflt_val, pk sql.NullString
		if err = rows.Scan(&count, &got, &LastName, &fieldNil, &dflt_val, &pk); err != nil {
			log.Fatalf("rows.Scan error: %v", err)
		}
		if expected[indexCount] != got {
			t.Errorf("expected = %s, got = %s", expected[indexCount], got)
		}
		indexCount++
	}
}
