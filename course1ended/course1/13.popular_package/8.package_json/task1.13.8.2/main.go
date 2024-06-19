package main

import (
	"encoding/json"
)

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}
type Comment struct {
	Text string `json:"text"`
}

// getUsersFromJSON которая принимает в качестве аргумента срез байт data и возвращает срез структур User и ошибку error.
// Функция getUsersFromJSON должна производить декодирование данных из формата JSON и возвращать срез структур User.
// Если произошла ошибка при декодировании, функция должна возвращать ошибку.
func getUsersFromJSON(data []byte) ([]User, error) {
	if len(data) < 1 {
		return nil, nil
	}
	var users []User
	err := json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

//func main() {
//	jsonData := []byte(`[
//		{
//			"name": "John",
//			"age": 30,
//			"comments": [
//				{"text": "Great post!"},
//				{"text": "I agree"}
//			]
//		},
//		{
//			"name": "Alice",
//			"age": 25,
//			"comments": [
//				{"text": "Nice article"}
//			]
//		}
//	]`)
//	users, err := getUsersFromJSON(jsonData)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//	for _, user := range users {
//		fmt.Println("Name:", user.Name)
//		fmt.Println("Age:", user.Age)
//		fmt.Println("Comments:")
//		for _, comment := range user.Comments {
//			fmt.Println("- ", comment.Text)
//		}
//		fmt.Println()
//	}
//}
