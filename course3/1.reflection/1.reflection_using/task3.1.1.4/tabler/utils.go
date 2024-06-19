package tabler

import (
	"reflect"
)

type Tabler interface {
	TableName() string
}

type StructInfo struct {
	Fields   []string
	Pointers []interface{}
}

func GetStructInfo(u interface{}, args ...func(*[]reflect.StructField)) StructInfo {
	val := reflect.ValueOf(u).Elem()
	var structFields []reflect.StructField

	for i := 0; i < val.NumField(); i++ {
		structFields = append(structFields, val.Type().Field(i))
	}

	for i := range args {
		if args[i] == nil {
			continue
		}
		args[i](&structFields)
	}

	var res StructInfo

	for _, field := range structFields {
		valueField := val.FieldByName(field.Name)
		res.Pointers = append(res.Pointers, valueField.Addr().Interface())
		res.Fields = append(res.Fields, field.Tag.Get("db"))
	}

	return res
}

func FilterByFields(fields ...int) func(fields *[]reflect.StructField) {
	requiredFields := make([]int, len(fields))
	for i, f := range fields {
		requiredFields[i] = f
	}
	return func(fs *[]reflect.StructField) {
		var res []reflect.StructField
		for i := range requiredFields {
			for j := range *fs {
				val := requiredFields[i] // переменные среза - ID, Username, Email = 1,2,4
				idx := 1 << j            // при сдвиге 1 << j индексы полей(j) = 0,1,2,3... получаемые значения = 1,2,4,8...
				if val&idx != 0 {        // true если хоть один бит 1:1,  false при - 1:0, 0:1, 0:0
					res = append(res, (*fs)[j])
					break
				}
			}
		}
		*fs = res
	}
}

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
