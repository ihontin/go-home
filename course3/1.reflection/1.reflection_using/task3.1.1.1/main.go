package main

import (
	"fmt"
	"reflect"
)

//Создай функцию, которая будет генерировать список указателей на поля структуры.

type User struct {
	ID       int    `db:"id" db_ops:"create"`
	Username string `db:"username" db_ops:"create,update"`
	Email    string `db:"email" db_ops:"create,update"`
	Address  string `db:"address" db_ops:"update"`
	Status   int    `db:"status" db_ops:"create,update"`
	Delete   string `db:"delete" db_ops:"delete"`
}

func SimpleGetFieldsPointers(u interface{}) []interface{} {
	var fiels = make([]interface{}, 0, 4)
	fmt.Println(u.(User))
	user := reflect.TypeOf(u)
	for i := 0; i < user.NumField(); i++ {
		field := user.Field(i)
		if tagNfield, _ := field.Tag.Lookup("db"); tagNfield == "delete" || tagNfield == "id" {
			continue
		}
		fiels = append(fiels, &field)
	}
	return fiels
}

func main() {
	user := User{
		ID:       1,
		Username: "JohnDoe",
		Email:    "johndoe@example.com",
		Address:  "123 Main St",
		Status:   1,
		Delete:   "yes",
	}
	pointers := SimpleGetFieldsPointers(user)
	for _, val := range pointers {
		fmt.Println(val.(*reflect.StructField).Tag.Get("db"))
	}
}
