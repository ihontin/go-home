package main

// User Определение структуры пользователя
type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

// Tabler Интерфейс Tabler содержит метод TableName который возвращает имя таблицы.
type Tabler interface {
	TableName() string
}

// SQLGenerator Интерфейс для генерации SQL-запросов
type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(table Tabler) string
}

// FakeDataGenerator Интерфейс для генерации фейковых данных
type FakeDataGenerator interface {
	GenerateFakeUser() User
}

//func main() {
//	sqlGenerator := &SQLiteGenerator{}
//	fakeDataGenerator := &GoFakeitGenerator{}
//
//	user := User{}
//	sql := sqlGenerator.CreateTableSQL(&user)
//	fmt.Println(sql)
//
//	for i := 0; i < 34; i++ {
//		fakeUser := fakeDataGenerator.GenerateFakeUser()
//		query := sqlGenerator.CreateInsertSQL(&fakeUser)
//		fmt.Println(query)
//	}
//}
