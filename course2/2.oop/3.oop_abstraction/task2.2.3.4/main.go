package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"reflect"
	"strings"
)

// User Определение структуры пользователя
type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

// Tabler Интерфейс Tabler содержит метод TableName который возвращает имя таблицы.
type Tabler interface {
	TableName() *User
}

func (u *User) TableName() *User {
	return u
}

type SQLiteGenerator struct {
	User
}

// SQLGenerator Интерфейс для генерации SQL-запросов
type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(table Tabler) string
}

func (s *SQLiteGenerator) CreateTableSQL(table Tabler) string {
	data := table.TableName()
	t := reflect.ValueOf(s.User).Type()
	create := "CREATE TABLE IF NOT EXISTS"
	FirstNameField := t.Field(1).Tag.Get("db_field")
	FirstNameType := t.Field(1).Tag.Get("db_type")
	LastNameField := t.Field(2).Tag.Get("db_field")
	LastNameType := t.Field(2).Tag.Get("db_type")
	EmailField := t.Field(3).Tag.Get("db_field")
	EmailType := t.Field(3).Tag.Get("db_type")
	idName := t.Field(0).Tag.Get("db_field")
	idVal := t.Field(0).Tag.Get("db_type")
	outRequest := fmt.Sprintf("%s %ss (%s %s, %s %s, %s %s, "+
		"%s %s);", create, strings.ToLower(reflect.TypeOf(*data).Name()), idName, idVal,
		FirstNameField, FirstNameType, LastNameField, LastNameType, EmailField, EmailType)
	fmt.Println(outRequest)
	return outRequest
}
func (s *SQLiteGenerator) CreateInsertSQL(table Tabler) string {
	data := table.TableName()
	t := reflect.ValueOf(s.User).Type()
	insert := "INSERT INTO"
	FirstNameField := t.Field(1).Tag.Get("db_field")
	LastNameField := t.Field(2).Tag.Get("db_field")
	EmailField := t.Field(3).Tag.Get("db_field")
	idName := t.Field(0).Tag.Get("db_field")
	outRequest := fmt.Sprintf("%s %ss (%s, %s, %s, %s) VALUES (%d,'%s','%s','%s');",
		insert, strings.ToLower(reflect.TypeOf(*data).Name()), idName, FirstNameField, LastNameField, EmailField,
		data.ID, data.FirstName, data.LastName, data.Email)
	return outRequest
}

//-----------------------------------------------------------------

type Migrator struct {
	db           *sql.DB
	sqlGenerator SQLGenerator
}

func NewMigrator(newDb *sql.DB, sqlGen SQLGenerator) *Migrator {
	return &Migrator{newDb, sqlGen}
}

func (m *Migrator) Migrate(tables ...Tabler) error {
	for _, table := range tables {
		_, _ = m.db.Exec(m.sqlGenerator.CreateTableSQL(table))
		//if err != nil {
		//	return err
		//}
	}
	return nil
}

// Основная функция
func main() {
	// Подключение к SQLite БД
	db, err := sql.Open("sqlite3", "file:my_database.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Создание мигратора с использованием вашего SQLGenerator
	YourSQLGeneratorInstance := &SQLiteGenerator{}
	migrator := NewMigrator(db, YourSQLGeneratorInstance)

	// Миграция таблицы User
	if err = migrator.Migrate(&User{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}
