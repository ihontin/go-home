package main

import (
	"container/list"
	"fmt"
)

type Car struct {
	LicensePlate string
}

type ParkingLot struct {
	space *list.List
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{space: list.New()}
}

func (p *ParkingLot) Park(c Car) {
	p.space.PushBack(c)
	fmt.Printf("Автомобиль {%s} припаркован.\n", c)
}

func (p *ParkingLot) Leave() {
	if p.space.Len() > 0 {
		leaveCar := p.space.Front()
		p.space.Remove(leaveCar)
		fmt.Printf("Автомобиль {%s} покинул парковку.\n", leaveCar.Value)
	} else {
		fmt.Println("Парковка пуста.")
	}
}

func main() {
	parkingLot := NewParkingLot()
	parkingLot.Park(Car{LicensePlate: "ABC-123"})
	parkingLot.Park(Car{LicensePlate: "XYZ-789"})
	parkingLot.Leave()
	parkingLot.Leave()
	parkingLot.Leave()
}
