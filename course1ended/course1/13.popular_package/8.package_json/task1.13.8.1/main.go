package main

import (
	"encoding/json"
	"fmt"
)

// Напиши функцию getJSON, которая принимает срез структур User и возвращает строку в формате JSON и ошибку.
func getJSON(data []User) (string, error) {
	if len(data) < 1 {
		return "[]", nil
	}
	users, err := json.Marshal(data)
	if err != nil {
		return "[]", err
	}
	return string(users), nil
}

//Структура User определена следующим образом:

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

//Структура Comment определена следующим образом:

type Comment struct {
	Text string `json:"text"`
}

func main() {
	var newUser = []User{
		{"Alex", 46, []Comment{
			{"what can i say"},
			{"everithing"},
		},
		},
		{"Susanna", 14, []Comment{
			{"Shiny star"},
			{"a cup of milk"},
		},
		},
	}
	marshaler, err := getJSON(newUser)
	if err != nil {
		fmt.Errorf("%v", err.Error())
	}
	fmt.Println(marshaler)
}
