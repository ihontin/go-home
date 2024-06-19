package main

import (
	"time"
)

type Order struct {
	ID         int
	CustomerID string
	Items      []string
	OrderDate  time.Time
}

type OrderOption func(*Order)

func WithCustomerID(s string) OrderOption {
	return func(order *Order) {
		order.CustomerID = s
	}
}
func WithItems(s []string) OrderOption {
	return func(order *Order) {
		order.Items = s
	}
}
func WithOrderDate(s time.Time) OrderOption {
	return func(order *Order) {
		order.OrderDate = s
	}
}
func NewOrder(id int, options ...OrderOption) *Order {
	newOrder := &Order{
		id, "", []string{}, time.Time{},
	}
	for _, option := range options {
		option(newOrder)
	}
	return newOrder
}

//func main() {
//	order := NewOrder(1,
//		WithCustomerID("123"),
//		WithItems([]string{"item1", "item2"}),
//		WithOrderDate(time.Now()))
//
//	fmt.Printf("Order: %+v\n", order)
//}
