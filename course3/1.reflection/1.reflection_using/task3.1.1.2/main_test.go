package main

import (
	"reflect"
	"testing"
)

func TestFilterByFields(t *testing.T) {
	u := User{
		ID:       1,
		Username: "john_doe",
		Email:    "john@example.com",
	}
	type args struct {
		fields []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"ok",
			args{fields: []int{ID, Username, Email}},
			[]string{"ID", "Username", "Email"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := reflect.TypeOf(u) // возвращает тип аргумента, для работы с тегами и данными аргумента
			var fiels = make([]reflect.StructField, 0, user.NumField())
			for i := 0; i < user.NumField(); i++ {
				field := user.Field(i)       // получает поле по индексу полей структуры
				fiels = append(fiels, field) // создаем срез всех полей
			}
			FilterByFields(tt.args.fields...)(&fiels)
			for i, val := range fiels {
				if !reflect.DeepEqual(val.Name, tt.want[i]) {
					t.Errorf("FilterByFields() = %v, want %v", val.Name, tt.want[i])
				}
			}
		})
	}
}

func TestFilterByTags(t *testing.T) {
	u := User{
		ID:       1,
		Username: "john_doe",
		Email:    "john@example.com",
	}
	type args struct {
		tags map[string]func(value string) bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"ok",
			args{tags: map[string]func(value string) bool{
				"db": func(value string) bool {
					values := []string{"id", "username", "email"}
					for _, v := range values {
						if v == value {
							return true
						}
					}
					return false
				},
			}},
			[]string{"ID", "Username", "Email"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := reflect.TypeOf(u) // возвращает тип аргумента, для работы с тегами и данными аргумента
			var fiels = make([]reflect.StructField, 0, user.NumField())
			for i := 0; i < user.NumField(); i++ {
				field := user.Field(i)       // получает поле по индексу полей структуры
				fiels = append(fiels, field) // создаем срез всех полей
			}
			FilterByTags(tt.args.tags)(&fiels)
			if len(fiels) != len(tt.want) {
				t.Errorf("FilterByTags() len = %v, want len = %v", len(fiels), len(tt.want))
				return
			}
			for i := range tt.want {
				if !reflect.DeepEqual(fiels[i].Name, tt.want[i]) {
					t.Errorf("FilterByTags() = %v, want = %v", fiels[i].Name, tt.want[i])
				}
			}
		})
	}
}

func TestGetFieldsPointers(t *testing.T) {
	u := User{
		ID:       1,
		Username: "john_doe",
		Email:    "john@example.com",
	}
	filter := func(fields *[]reflect.StructField) {
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
	type args struct {
		u    interface{}
		args []func(*[]reflect.StructField)
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"ok",
			args{u: &u, args: []func(*[]reflect.StructField){filter}},
			[]string{"", "john_doe", "john@example.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pointers := GetFieldsPointers(tt.args.u, tt.args.args...)
			for i := range pointers {
				switch pointers[i].(type) {
				case *string:
					if *pointers[i].(*string) != tt.want[i] {
						t.Errorf("GetFieldsPointers() = %v, want %v", *pointers[i].(*string), tt.want[i])
					}
				}
			}

		})
	}
}
