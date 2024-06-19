package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
	"time"
)

type User struct {
	Name string
	Age  int
}

// getUsers () должна возвращать срез пользователей.
// Каждый пользователь должен иметь поля Name (имя) и Age (возраст).
// Для генерации случайных данных о пользователях используй пакет gofakeit.
func getUsers() []User {
	rand.Seed(time.Now().UnixNano())
	var newUser []User
	for i := 0; i < 10; i++ {
		name, age := func() (string, int) {
			return gofakeit.FirstName(), rand.Intn(42) + 18
		}()
		newUser = append(newUser, User{name, age})
	}
	return newUser
}

// preparePrint ([]User) должна принимать срез пользователей и возвращать строку,
// содержащую информацию о каждом пользователе в формате “Имя: %s, Возраст: %d”.
// Каждый пользователь должен быть представлен на отдельной строке.
func preparePrint(u []User) string {
	var allUsers string
	for _, user := range u {
		allUsers += fmt.Sprintf("Имя: %s, Возраст: %d\n", user.Name, user.Age)
	}
	return allUsers
}

//func main() {
//	users := getUsers() // Получаем срез 10 пользователей
//	result := preparePrint(users)
//	fmt.Println(result)
//}
