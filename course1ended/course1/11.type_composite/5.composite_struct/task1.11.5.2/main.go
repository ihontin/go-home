package main

import "fmt"

// Пример кода
type Animal struct {
	Type string
	Name string
	Age  int
}

// getAnimals () должна возвращать срез (массив) животных типа Animal.
func getAnimals() []Animal {
	var threeAni = []Animal{
		{"dog", "Barbos", 5},
		{"turtle", "Speedy", 6},
		{"eagle", "Tromb", 7},
	}
	return threeAni
}

// preparePrint ([]Animal) string должна принимать срез животных и возвращать строку,
// содержащую информацию о каждом животном в формате “Тип: %s, Имя: %s, Возраст: %d”,
// где %s - тип животного, %s - имя животного, %d - возраст животного.
func preparePrint(animals []Animal) string {
	// Ваш код для форматирования информации о животных
	var allUsers string
	for _, an := range animals {
		allUsers += fmt.Sprintf("Тип: %s, Имя: %s, Возраст: %d\n", an.Type, an.Name, an.Age)
	}
	return allUsers
}

//func main() {
//	users := getAnimals() // Получаем срез
//	result := preparePrint(users)
//	fmt.Println(result)
//}
