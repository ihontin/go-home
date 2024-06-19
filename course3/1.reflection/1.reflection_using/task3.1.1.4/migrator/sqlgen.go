package migrator

import (
	"fmt"
	"reflect"
	"studentgit.kata.academy/Alkolex/go-kata/course3/1.reflection/1.reflection_using/task3.1.1.4/tabler"
)

type SQLGenerator interface {
	CreateTableSQL(table tabler.Tabler) string
}

type SQLiteGenerator struct{}

func (sg *SQLiteGenerator) CreateTableSQL(table tabler.Tabler) string {
	// Создание SQL-запроса для создания таблицы
	query := "CREATE TABLE IF NOT EXISTS " + fmt.Sprintf("%s (", table.TableName())
	val := reflect.TypeOf(table).Elem()
	for i := 0; i < val.NumField(); i++ {
		query += val.Field(i).Tag.Get("db") + " " + val.Field(i).Tag.Get("db_type")
		if i+1 == val.NumField() {
			query += ");" // Закрытие скобок при последнем добавлении полей
		} else {
			query += ", "
		}
	}
	return query
}
