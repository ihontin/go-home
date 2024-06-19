package main

import (
	"encoding/json"
	"fmt"
	"github.com/mailru/easyjson"
	"log"
	"studentgit.kata.academy/Alkolex/go-kata/course2/5.optimization/homework/practice/models"
)

// easyjson -all models.go
//команда терминала для генерации файла easyjson

func main() {
	myUser := models.Users{"Natasha", 15, false}
	byteUser, err := json.Marshal(myUser) // привычный Marshal
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteUser))
	var makeUser models.Users
	err = json.Unmarshal(byteUser, &makeUser) // привычный Unmarshal
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(makeUser)
	easyjsonUser, err := easyjson.Marshal(myUser) // easyjson Marshal
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(easyjsonUser))
	var easyUser models.Users
	err = easyjson.Unmarshal(easyjsonUser, &easyUser) // easyjson Unmarshal
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(easyUser)
}
