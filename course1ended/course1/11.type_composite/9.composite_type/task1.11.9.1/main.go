package main

import "fmt"

// —  Live() - метод принадлежащий интерфейсу OrganicWorld
type OrganicWorld interface {
	Live()
}

// —  Replicate() NonCellular - метод принадлежащий интерфейсу NonCellular, возвращает новый объект интерфейсного типа NonCellular
type NonCellular interface {
	OrganicWorld
	Replicate() NonCellular
	Virus
}

// —  Infect() - метод принадлежащий интерфейсу Virus
type Virus interface {
	Infect()
}

func (v *InfluenzaVirus) Live() {
}
func (v *InfluenzaVirus) Replicate() NonCellular {
	return v
}
func (v *InfluenzaVirus) Infect() {
}

// —  Grow() - метод принадлежащий интерфейсу Cell
// —  Divide() Cell - метод принадлежащий интерфейсу Cell, возвращает новый объект интерфейсного типа Cell
type Cell interface {
	OrganicWorld
	Grow()
	Divide() Cell
}

// — CloneGenome() - метод принадлежащий интерфейсу Eukaryote
type Eukaryote interface {
	Cell
	CloneGenome()
}

func (aE *AnimalEukaryote) Divide() Cell {
	return aE
}
func (aE *AnimalEukaryote) Grow() {
}
func (aE *AnimalEukaryote) Live() {
}
func (aE *AnimalEukaryote) CloneGenome() {
}

// —  Move() - метод принадлежащий интерфейсу Animal
// —  Eat() - метод принадлежащий интерфейсу Animal
type Animal interface {
	Eukaryote
	Move()
	Eat()
}

func (aCat *AnimalCat) Divide() Cell {
	return aCat
}
func (aCat *AnimalCat) Grow() {
}
func (aCat *AnimalCat) CloneGenome() {
}
func (aCat *AnimalCat) Live() {
}
func (aCat *AnimalCat) Move() {
}
func (aCat *AnimalCat) Eat() {
}

// — ProduceToxins() - метод принадлежащий интерфейсу Prokaryote
type Prokaryote interface {
	Cell
	ProduceToxins()
}

type AnimalCat struct {
	AnimalEukaryote
}
type AnimalEukaryote struct {
	Cell
}
type AnimalCell struct {
	Cell
}
type InfluenzaVirus struct {
	Virus
}

func main() {

	var cell Cell = &AnimalCat{}
	cell.Grow()
	newCell := cell.Divide()
	fmt.Println(newCell)

	var nonCell NonCellular = &InfluenzaVirus{}
	nonCell.Infect()
	newNonCell := nonCell.Replicate()
	fmt.Println(newNonCell)
}
