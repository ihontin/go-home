package main

import "fmt"

// Функция должна вернуть строку с информацией о пользователе в формате “Имя: {name}, возраст: {age}, города: {cities}”.
func UserInfo(name string, age int, cities ...string) string {
	var strCities string
	for i, city := range cities {
		if i > 0 {
			strCities += ", " + city
		} else {
			strCities += city
		}
	}
	return fmt.Sprintf("Имя: %s, возраст: %d, города: %v", name, age, strCities)
}
func main() {

	result := UserInfo("John", 21, "Moscow", "Saint Petersburg")
	fmt.Println(result)
	//Имя: John, возраст: 21, города: Moscow, Saint Petersburg

	result = UserInfo("Alex", 34)
	fmt.Println(result)
	//Имя: Alex, возраст: 34, города:
}
