package main

type Dish struct {
	Name  string
	Price float64
}

type Order struct {
	Dishes []Dish
	Total  float64
}

// AddDish должна добавлять блюдо в заказ

func (order *Order) AddDish(dish Dish) {
	order.Dishes = append(order.Dishes, dish)
}

// RemoveDish должна удалять блюдо из заказа
func (order *Order) RemoveDish(dish Dish) {
	for i, ordDish := range order.Dishes {
		if ordDish.Name == dish.Name {
			order.Dishes = append(order.Dishes[:i], order.Dishes[i+1:]...)
		}
	}
}

// CalculateTotal должна вычислять общую стоимость заказа. Общая стоимость должна сохраняться в поле order.Total.
func (order *Order) CalculateTotal() {
	order.Total = 0.0
	for _, dish := range order.Dishes {
		order.Total += dish.Price
	}
}

//func main() {
//	order := Order{}
//	dish1 := Dish{Name: "Pizza", Price: 10.99}
//	dish2 := Dish{Name: "Burger", Price: 5.99}
//
//	order.AddDish(dish1)
//	order.AddDish(dish2)
//
//	order.CalculateTotal()
//	fmt.Println("Total:", order.Total)
//
//	order.RemoveDish(dish1)
//
//	order.CalculateTotal()
//	fmt.Println("Total:", order.Total)
//}
