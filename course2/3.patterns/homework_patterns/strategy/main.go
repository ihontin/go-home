package main

import "fmt"

//В игровом движке может быть несколько различных стратегий для обработки физики объектов,
//их взаимодействия и искусственного интеллекта. Паттерн стратегия позволяет легко добавлять новые стратегии
//и изменять их взаимодействие без изменения основного кода движка.
// Интерфейс стратегии

//type PhysicsStrategy interface {
//	UpdatePosition()
//}
//
//// Конкретная стратегия для обработки физики объектов
//type SimplePhysicsStrategy struct{}
//
//func (s *SimplePhysicsStrategy) UpdatePosition() {
//	fmt.Println("Обновление позиции объекта по простой физике")
//}
//
//// Конкретная стратегия для обработки физики объектов с учетом силы тяжести
//type GravityPhysicsStrategy struct{}
//
//func (s *GravityPhysicsStrategy) UpdatePosition() {
//	fmt.Println("Обновление позиции объекта с учетом силы тяжести")
//}
//
//// Контекст
//type GameObject struct {
//	physicsStrategy PhysicsStrategy
//}
//
//func (g *GameObject) Update() {
//	g.physicsStrategy.UpdatePosition()
//}
//
//func main() {
//	// Создание объекта и выбор стратегии
//	obj := &GameObject{physicsStrategy: &SimplePhysicsStrategy{}}
//
//	// Обновление позиции объекта
//	obj.Update()
//
//	// Изменение стратегии
//	obj.physicsStrategy = &GravityPhysicsStrategy{}
//
//	// Обновление позиции объекта с учетом силы тяжести
//	obj.Update()
//}

//Пример 2. Разработка системы оптимизации ресурсов
//В системе оптимизации ресурсов может быть несколько стратегий для управления распределением и использованием ресурсов,
//таких как память, процессорное время и сетевые ресурсы. Паттерн стратегия позволяет динамически выбирать
//и изменять стратегии в зависимости от текущих требований и условий.
// Интерфейс стратегии
//type ResourceStrategy interface {
//	Allocate()
//}
//
//// Конкретная стратегия для управления памятью
//type MemoryStrategy struct{}
//
//func (s *MemoryStrategy) Allocate() {
//	fmt.Println("Выделение памяти для ресурса")
//}
//
//// Конкретная стратегия для управления процессорным временем
//type CPUStrategy struct{}
//
//func (s *CPUStrategy) Allocate() {
//	fmt.Println("Выделение процессорного времени для ресурса")
//}
//
//// Контекст
//type ResourceManager struct {
//	resourceStrategy ResourceStrategy
//}
//
//func (r *ResourceManager) AllocateResource() {
//	r.resourceStrategy.Allocate()
//}
//
//func main() {
//	// Создание менеджера ресурсов и выбор стратегии
//	manager := &ResourceManager{resourceStrategy: &MemoryStrategy{}}
//
//	// Выделение памяти для ресурса
//	manager.AllocateResource()
//
//	// Изменение стратегии
//	manager.resourceStrategy = &CPUStrategy{}
//
//	// Выделение процессорного времени для ресурса
//	manager.AllocateResource()
//}
//
//…

//Приведенные примеры демонстрируют, как паттерн стратегия может быть использован для управления различными
//алгоритмами в реальных проектах, обеспечивая гибкость и расширяемость системы.

type haveFun interface {
	MakeMeLafe()
}

type HeatToPain struct{}

func (h *HeatToPain) MakeMeLafe() {
	fmt.Println("Ты ударился - это смешно")
}

type GoodJoke struct{}

func (g *GoodJoke) MakeMeLafe() {
	fmt.Println("Хорошая шутка, но не смешная")
}

type СhooseJoke struct {
	haveFun haveFun
}

func (c *СhooseJoke) makeFun() {
	c.haveFun.MakeMeLafe()
}

func main() {
	a := &СhooseJoke{haveFun: &HeatToPain{}}

	a.makeFun()

	a.haveFun = &GoodJoke{}

	a.makeFun()
}
