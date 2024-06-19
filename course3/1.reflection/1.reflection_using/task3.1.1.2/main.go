package main

import (
	"fmt"
	"reflect"
)

const (
	ID       = 1 << iota // 1 << 0 == 1
	Username             // 1 << 1 == 2
	Email                // 1 << 2 == 4
	Address              // 1 << 3 == 8
	Status               // 1 << 4 == 16
)

type User struct {
	ID       int    `db:"id" db_ops:"create"`
	Username string `db:"username" db_ops:"create,update"`
	Email    string `db:"email" db_ops:"create,update"`
	Address  string `db:"address" db_ops:"update"`
	Status   int    `db:"status" db_ops:"create,update"`
	Delete   string `db:"delete" db_ops:"delete"`
}

// GetFieldsPointers возвращает срез значений полей экземпляра структуры,
// u - экземпляра структуры, args - опции для маппинга полей структуры
func GetFieldsPointers(u interface{}, args ...func(*[]reflect.StructField)) []interface{} {
	user := reflect.TypeOf(*u.(*User))   // возвращает тип аргумента, для работы с тегами и данными аргумента
	userVal := reflect.ValueOf(u).Elem() // возвращает значение аргумента для взаимодействия с ним
	var fiels = make([]reflect.StructField, 0, user.NumField())
	//query := "CREATE TABLE IF NOT EXISTS " + fmt.Sprint("users (")
	for i := 0; i < user.NumField(); i++ {
		//query += user.Field(i).Tag.Get("db") + " " + user.Field(i).Tag.Get("db_ops")
		//if i+1 == user.NumField() {
		//	query += ");"
		//} else {
		//	query += ", "
		//}
		fmt.Println(userVal.Addr().Interface())
		field := user.Field(i)       // получает поле db_opsпо индексу полей структуры
		fiels = append(fiels, field) // создаем срез всех полей
	}
	//fmt.Println(query)
	for i := range args {
		if args[i] == nil {
			continue
		}
		args[i](&fiels) // опции для маппинга полей возвращают отфильтрованный срез полей
	}
	var fiInter = make([]interface{}, 0, len(fiels))
	for _, val := range fiels {
		var fildV = userVal.Field(val.Index[0]) // достаем значение по индексу поля из отфильтрованного среза
		if fildV.IsZero() {                     // если значение отсутствует не сохраняем поле
			continue
		}
		// получаем изначальный тип переменной и добавляем в срез
		if resValStr, okS := fildV.Interface().(string); okS {
			fiInter = append(fiInter, &resValStr)
		} else if resValInt, okI := fildV.Interface().(int); okI {
			fiInter = append(fiInter, &resValInt)
		}
	}
	return fiInter
}

func main() {
	user := User{
		ID:       1,
		Username: "john_doe",
		Email:    "john@example.com",
	}
	// опция выбора полей переменных
	filter1 := func(fields *[]reflect.StructField) {
		var res []reflect.StructField
		requiredFields := []int{ID, Username, Email}
		for i := range requiredFields {
			for j := range *fields {
				val := requiredFields[i] // переменные среза - ID, Username, Email = 1,2,4
				idx := 1 << j            // при сдвиге 1 << j индексы полей(j) = 0,1,2,3... получаемые значения = 1,2,4,8...
				if val&idx != 0 {        // true если хоть один бит 1:1,  false при - 1:0, 0:1, 0:0
					res = append(res, (*fields)[j])
					break
				}
			}
		}
		*fields = res
	}
	// опция выбора полей переменных
	filter2 := func(fields *[]reflect.StructField) {
		var res []reflect.StructField

		for i := range *fields {
			if (*fields)[i].Tag.Get("db_ops") != "create" {
				res = append(res, (*fields)[i])
			}
		}

		*fields = res
	}
	pointers := GetFieldsPointers(&user, filter1, filter2)
	for _, pointer := range pointers {
		switch v := pointer.(type) {
		case *int:
			fmt.Printf("%T %v: %v\n", pointer, pointer, *v)
		case *string:
			fmt.Printf("%T %v: %v\n", pointer, pointer, *v)
		}
	}
	fmt.Println()

	pointers = GetFieldsPointers(&user, FilterByFields(ID, Username, Email))
	fmt.Println("FilterByFields(ID, Username, Email)")

	for _, pointer := range pointers {
		switch v := pointer.(type) {
		case *int:
			fmt.Printf("%T %v: %v\n", pointer, pointer, *v)
		case *string:
			fmt.Printf("%T %v: %v\n", pointer, pointer, *v)
		}
	}
	// карта для сопоставления значений тегов по ключу и значению в опции выбора полей переменных
	filterTag := map[string]func(value string) bool{
		"db": func(value string) bool {
			values := []string{"username", "address", "status"}
			for _, v := range values {
				if v == value {
					return true
				}
			}
			return false
		},
	}

	fmt.Println()

	pointers = GetFieldsPointers(&user, FilterByTags(filterTag))
	fmt.Println("FilterByTags(filterTag)")

	for _, pointer := range pointers {
		switch v := pointer.(type) {
		case *int:
			fmt.Printf("%T %v: %v\n", pointer, pointer, *v)
		case *string:
			fmt.Printf("%T %v: %v\n", pointer, pointer, *v)
		}
	}
}

// FilterByFields опция выбора полей переменных
func FilterByFields(fields ...int) func(fields *[]reflect.StructField) {
	requiredFields := make([]int, len(fields))
	for i, f := range fields {
		requiredFields[i] = f
	}
	return func(fields *[]reflect.StructField) {
		var res []reflect.StructField
		for i := range requiredFields {
			for j := range *fields {
				val := requiredFields[i] // переменные среза - ID, Username, Email = 1,2,4
				idx := 1 << j            // при сдвиге 1 << j индексы полей(j) = 0,1,2,3... получаемые значения = 1,2,4,8...
				if val&idx != 0 {        // true если хоть один бит 1:1,  false при - 1:0, 0:1, 0:0
					res = append(res, (*fields)[j])
					break
				}
			}
		}
		*fields = res
	}
}

// FilterByTags опция выбора полей переменных
func FilterByTags(tags map[string]func(value string) bool) func(fields *[]reflect.StructField) {
	return func(fields *[]reflect.StructField) {
		var res []reflect.StructField
		// key - название и value - значение тега поля
		for key, value := range tags {
			for i := range *fields {
				// находим значение тега по ключу `db` проверяем его на соответствие в
				//value(data) = func(value string) bool
				if data, ok := (*fields)[i].Tag.Lookup(key); ok && value(data) {
					res = append(res, (*fields)[i])
				}
			}
		}
		*fields = res
	}
}
