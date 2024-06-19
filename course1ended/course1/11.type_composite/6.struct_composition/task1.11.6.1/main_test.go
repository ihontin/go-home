package main

import (
	"testing"
)

func addNamePrise(order *Order) {
	dish3 := Dish{Name: "Tutu", Price: 2.66}
	dish4 := Dish{Name: "Gaga", Price: 3.33}
	dish5 := Dish{Name: "Lolo", Price: 4.88}
	dish6 := Dish{Name: "Fifi", Price: 7.77}

	order.AddDish(dish3)
	order.AddDish(dish4)
	order.AddDish(dish5)
	order.AddDish(dish6)
}

func TestAddDish(t *testing.T) {
	order := Order{}
	addNamePrise(&order)
	expected := "Lolo"
	expectedPrice := 4.88
	if order.Dishes[2].Name != expected || order.Dishes[2].Price != expectedPrice {
		t.Errorf("expected name = %s, expected price = %f, got name = %s, got prise = %f",
			expected, expectedPrice, order.Dishes[2].Name, order.Dishes[2].Price)
	}
}
func TestRemoveDish(t *testing.T) {
	order := Order{}
	addNamePrise(&order)
	order.RemoveDish(order.Dishes[2])
	expected := "Fifi"
	expectedPrice := 7.77
	if order.Dishes[2].Name != expected || order.Dishes[2].Price != expectedPrice {
		t.Errorf("expected name = %s, expected price = %f, got name = %s, got prise = %f",
			expected, expectedPrice, order.Dishes[2].Name, order.Dishes[2].Price)
	}
}
func TestCalculateTotal(t *testing.T) {
	order := Order{}
	addNamePrise(&order)
	order.CalculateTotal()
	expected := 4.88 + 2.66 + 3.33 + 7.77
	if order.Total != expected {
		t.Errorf("expected total = %f, got = %f",
			expected, order.Total)
	}
}
