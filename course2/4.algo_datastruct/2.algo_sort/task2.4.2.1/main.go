package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

type Product struct {
	Name      string
	Price     float64
	CreatedAt time.Time
	Count     int
}

func (p Product) String() string {
	return fmt.Sprintf("Name: %s, Price: %f, Count: %v", p.Name, p.Price, p.Count)
}

func generateProducts(n int) []Product {
	gofakeit.Seed(time.Now().UnixNano())
	products := make([]Product, n)
	for i := range products {
		products[i] = Product{
			Name:      gofakeit.Word(),
			Price:     gofakeit.Price(1.0, 100.0),
			CreatedAt: gofakeit.Date(),
			Count:     gofakeit.Number(1, 100),
		}
	}
	return products
}

//func main() {
//	products := generateProducts(10)
//
//	fmt.Println("Исходный список:")
//	fmt.Println(products)
//
//	// Сортировка продуктов по цене
//	sort.Sort(ByPrice(products))
//	fmt.Println("\nОтсортировано по цене:")
//	fmt.Println(products)
//
//	// Сортировка продуктов по дате создания
//	sort.Sort(ByCreatedAt(products))
//	fmt.Println("\nОтсортировано по дате создания:")
//	fmt.Println(products)
//
//	// Сортировка продуктов по количеству
//	sort.Sort(ByCount(products))
//	fmt.Println("\nОтсортировано по количеству:")
//	fmt.Println(products)
//}

type ByPrice []Product

func (bp ByPrice) Len() int           { return len(bp) }                   // Возвращает количество элементов в коллекции
func (bp ByPrice) Less(i, j int) bool { return bp[i].Price < bp[j].Price } // сравнивает элементы с индексами i, j
func (bp ByPrice) Swap(i, j int)      { bp[i], bp[j] = bp[j], bp[i] }      // меняет местами элементы с индексами i, j

type ByCreatedAt []Product

func (bp ByCreatedAt) Len() int           { return len(bp) }
func (bp ByCreatedAt) Less(i, j int) bool { return bp[i].CreatedAt.Unix() < bp[j].CreatedAt.Unix() }
func (bp ByCreatedAt) Swap(i, j int)      { bp[i], bp[j] = bp[j], bp[i] }

type ByCount []Product

func (bp ByCount) Len() int           { return len(bp) }
func (bp ByCount) Less(i, j int) bool { return bp[i].Count < bp[j].Count }
func (bp ByCount) Swap(i, j int)      { bp[i], bp[j] = bp[j], bp[i] }
