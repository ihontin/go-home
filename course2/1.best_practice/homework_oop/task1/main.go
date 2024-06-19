package main

import "fmt"

//Пример 1: Абстракция
//Абстракция в Golang достигается с помощью определения интерфейсов. Рассмотрим пример,
//где создается интерфейс Shape для абстрагирования общих характеристик геометрических фигур:

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func main() {
	var shape Shape

	shape = Rectangle{width: 5, height: 3}
	fmt.Println("Прямоугольник:")
	fmt.Println("Площадь:", shape.Area())
	fmt.Println("Периметр:", shape.Perimeter())

	shape = Circle{radius: 4}
	fmt.Println("Круг:")
	fmt.Println("Площадь:", shape.Area())
	fmt.Println("Периметр:", shape.Perimeter())
}

//Пример 2: Наследование
//Наследование в Golang реализуется с помощью встраивания структур или интерфейсов в другие структуры.
//	Рассмотрим пример, где создается иерархия классов Animal и Dog, где Dog наследует свойства и методы от Animal:

//type Animal struct {
//	name string
//}
//
//func (a Animal) Speak() {
//	fmt.Println("Я животное и умею издавать звуки.")
//}
//
//type Dog struct {
//	Animal
//	breed string
//}
//
//func (d Dog) Speak() {
//	fmt.Println("Я собака и говорю 'Гав-гав!'")
//}
//
//func main() {
//	animal := Animal{name: "Животное"}
//	animal.Speak()
//
//	dog := Dog{Animal: Animal{name: "Собака"}, breed: "Лабрадор"}
//	dog.Speak()
//}

//Пример 3: Полиморфизм
//Полиморфизм в Golang достигается через интерфейсы. Рассмотрим пример, где создается интерфейс Player и две
//структуры FootballPlayer и BasketballPlayer, которые реализуют этот интерфейс:

//type Player interface {
//	Play()
//}
//
//type FootballPlayer struct {
//	name string
//}
//
//func (f FootballPlayer) Play() {
//	fmt.Println("Футболист", f.name, "играет в футбол.")
//}
//
//type BasketballPlayer struct {
//	name string
//}
//
//func (b BasketballPlayer) Play() {
//	fmt.Println("Баскетболист", b.name, "играет в баскетбол.")
//}
//
//func main() {
//	players := []Player{
//		FootballPlayer{name: "Роналду"},
//		BasketballPlayer{name: "Джордан"},
//	}
//
//	for _, player := range players {
//		player.Play()
//	}
//}

//Пример 4: Композиция
//Композиция в Golang достигается путем встраивания одной структуры в другую. Рассмотрим пример,
//где создается структура Car, которая состоит из структуры Engine и структуры Wheel:

//type Engine struct {
//power int
//}
//
//type Wheel struct {
//size int
//}
//
//type Car struct {
//engine Engine
//wheel  Wheel
//}
//
//func main() {
//car := Car{
//engine: Engine{power: 200},
//wheel:  Wheel{size: 18},
//}
//
//fmt.Println("Мощность двигателя:", car.engine.power)
//fmt.Println("Размер колес:", car.wheel.size)
//}

//Пример 5: Инкапсуляция
//Инкапсуляция в Golang достигается с помощью использования публичных и приватных полей и методов.
//	Рассмотрим пример, где создается структура Person с публичным полем Name и приватным полем age:

//type Person struct {
//Name string
//age  int
//}
//
//func (p Person) GetAge() int {
//return p.age
//}
//
//func main() {
//person := Person{Name: "Иван", age: 25}
//
//fmt.Println("Имя:", person.Name)
//fmt.Println("Возраст:", person.GetAge())
//}

//В данных примерах показано применение основных принципов ООП на языке Golang. Абстракция, наследование,
//полиморфизм, композиция и инкапсуляция являются важными концепциями, которые помогают создавать модульный,
//гибкий и расширяемый код.
