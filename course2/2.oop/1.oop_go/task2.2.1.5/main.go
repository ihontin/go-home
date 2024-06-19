package main

import "fmt"

// Mover Все типы, которые реализуют этот интерфейс, должны предоставить определение для каждого метода.
type Mover interface {
	Move() string
	Speed() int
	MaxSpeed() int
	MinSpeed() int
}

// BaseMover родительская структура хранит скорость движения
type BaseMover struct {
	speed int
}

// создание методов родительской структуры BaseMover, имплементирующая интерфейс Mover

func (b BaseMover) Move() string {
	return fmt.Sprintf("Base mover - Moving at speed: %d", b.speed)
}
func (b BaseMover) Speed() int {
	return b.speed
}
func (b BaseMover) MaxSpeed() int {
	return 120
}
func (b BaseMover) MinSpeed() int {
	return 10
}

// FastMover поля и методы встроенной структуры BaseMover становятся доступными
type FastMover struct {
	BaseMover
}

// Move - метод структуры FastMover заменит метод родительской структуры BaseMover
func (f FastMover) Move() string {
	return fmt.Sprintf("Fast mover! Moving at speed: %d", f.speed)
}

// SlowMover поля и методы встроенной структуры BaseMover становятся доступными
type SlowMover struct {
	BaseMover
}

// Move - метод структуры SlowMover заменит метод родительской структуры BaseMover
func (s SlowMover) Move() string {
	return fmt.Sprintf("Slow mover... Moving at speed: %d", s.speed)
}

func main() {
	var movers []Mover              // слайс муверов
	fm := FastMover{BaseMover{100}} // fm экземпляр структуры дочерней FastMover
	sm := SlowMover{BaseMover{10}}  // fm экземпляр структуры дочерней SlowMover
	// слайс movers может хранить в себе объекты не имплементирующие напрямую интерфейс Mover, но наследующихся от BaseMover
	movers = append(movers, fm, sm)
	for _, mover := range movers {
		fmt.Println(mover.Move())
		fmt.Println("Maximum speed:", mover.MaxSpeed())
		fmt.Println("Minimum speed:", mover.MinSpeed())
	}
}

//Output:
//
//Fast mover! Moving at speed: 100
//Maximum speed: 120
//Minimum speed: 10
//Slow mover... Moving at speed: 10
//Maximum speed: 120
//Minimum speed: 10
