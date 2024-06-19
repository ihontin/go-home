package main

import "fmt"

func main() {
	var name string
	var age int
	var city string

	fmt.Print("Введите ваше имя: ")
	fmt.Scanln(&name)
	fmt.Print("Введите ваш возраст: ")
	fmt.Scanln(&age)
	fmt.Print("Введите ваш город: ")
	fmt.Scanln(&city)

	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age)
	fmt.Println("Город:", city)

	//Имя: Petr
	//Возраст: 34
	//Город: Moscow
}
