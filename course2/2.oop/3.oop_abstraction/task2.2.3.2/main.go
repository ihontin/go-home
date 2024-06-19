package main

type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

type SQLiteGenerator struct{}

type Tabler interface {
	TableName() string
}

func (s *SQLiteGenerator) CreateTableSQL(table Tabler) string {
	panic("implement me")
}

type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(table Tabler) string
}

func (s *SQLiteGenerator) CreateInsertSQL(table Tabler) string {
	panic("implement me")
}

type FakeDataGenerator interface {
	GenerateFakeUser() User
}
type GoFakeitGenerator struct{}

func (g *GoFakeitGenerator) TableName() string {
	panic("implement me")
}
func (g *GoFakeitGenerator) GenerateFakeUser() User {
	panic("implement me")
}
