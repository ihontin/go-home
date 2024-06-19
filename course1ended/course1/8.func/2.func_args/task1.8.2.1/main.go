package main

import "fmt"

//	type AgeWeight struct {
//		age    int
//		weight int
//	}
//
//	func UserInfo(name, city, phone string, AgeWeight AgeWeight) string {
//		var outString string
//		if AgeWeight.age != 0 && AgeWeight.weight != 0 {
//			outString = fmt.Sprintf("Имя: %s, Город: %s, Телефон: %s, Возраст: %d, Вес: %d",
//				name, city, phone, AgeWeight.age, AgeWeight.weight)
//		} else if AgeWeight.age != 0 && AgeWeight.weight == 0 {
//			outString = fmt.Sprintf("Имя: %s, Город: %s, Телефон: %s, Возраст: %d",
//				name, city, phone, AgeWeight.age)
//		} else if AgeWeight.age == 0 && AgeWeight.weight != 0 {
//			outString = fmt.Sprintf("Имя: %s, Город: %s, Телефон: %s, Вес: %d",
//				name, city, phone, AgeWeight.weight)
//		} else if AgeWeight.age == 0 && AgeWeight.weight == 0 {
//			outString = fmt.Sprintf("Имя: %s, Город: %s, Телефон: %s", name, city, phone)
//		} else {
//			return "Error"
//		}
//
//		return outString
//	}
func UserInfo(name, city, phone string, age, weight int) string {
	var outString string
	if age != 0 && weight != 0 {
		outString = fmt.Sprintf("Имя: %s, Город: %s, Телефон: %s, Возраст: %d, Вес: %d",
			name, city, phone, age, weight)
	} else if age != 0 && weight == 0 {
		outString = fmt.Sprintf("Имя: %s, Город: %s, Телефон: %s, Возраст: %d",
			name, city, phone, age)
	} else if age == 0 && weight != 0 {
		outString = fmt.Sprintf("Имя: %s, Город: %s, Телефон: %s, Вес: %d",
			name, city, phone, weight)
	} else if age == 0 && weight == 0 {
		outString = fmt.Sprintf("Имя: %s, Город: %s, Телефон: %s", name, city, phone)
	} else {
		return "Error"
	}

	return outString
}
func main() {
	//instansAW := AgeWeight{
	//	age:    12,
	//	weight: 87,
	//}
	fmt.Println(UserInfo("Lalalu", "Luxemburg", "99-07-0990", 0, 0))
}
