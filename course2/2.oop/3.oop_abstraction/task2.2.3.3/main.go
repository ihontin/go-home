package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
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
	gfg GoFakeitGenerator
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

// FakeDataGenerator Интерфейс для генерации фейковых данных
type FakeDataGenerator interface {
	GenerateFakeUser() User
}

type GoFakeitGenerator struct{}

func (g *GoFakeitGenerator) GenerateFakeUser() User {
	idGen := gofakeit.Number(1000, 10000)
	FirstNameGen := gofakeit.FirstName()
	LastNameGen := gofakeit.LastName()
	EmailGen := gofakeit.Email()
	return User{ID: idGen, FirstName: FirstNameGen, LastName: LastNameGen, Email: EmailGen}
}

func main() {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}
	user := User{}
	sql := sqlGenerator.CreateTableSQL(&user)
	fmt.Println(sql)

	for i := 0; i < 34; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		fmt.Println(query)
	}
	queries := GenerateUserInserts(34)
	for _, query := range queries {
		fmt.Println(query)
	}
}

func GenerateUserInserts(n int) []string {
	outList := make([]string, 0, n)
	sqlGenerator := &SQLiteGenerator{}
	for i := 0; i < n; i++ {
		fakeUser := sqlGenerator.gfg.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		outList = append(outList, query)
	}
	return outList
}
