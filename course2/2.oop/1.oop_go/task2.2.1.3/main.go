package main

import (
	"fmt"
)

type Order interface {
	AddItem(item string, quantity int) error
	RemoveItem(item string) error
	GetOrderDetails() map[string]int
}
type DineInOrder struct {
	orderDetails map[string]int
}

func (d *DineInOrder) AddItem(item string, quantity int) error {
	if quantity < 1 {
		return fmt.Errorf("wrong quantity: %v", quantity)
	}
	d.orderDetails[item] = quantity
	fmt.Println(d.orderDetails)
	return nil
}
func (d *DineInOrder) RemoveItem(item string) error {
	if _, ok := d.orderDetails[item]; !ok {
		return fmt.Errorf("wrong item: %v", item)
	}
	delete(d.orderDetails, item)
	fmt.Println(d.orderDetails)
	return nil
}
func (d *DineInOrder) GetOrderDetails() map[string]int {
	return d.orderDetails
}

type TakeAwayOrder struct {
	orderDetails map[string]int
}

func (d *TakeAwayOrder) AddItem(item string, quantity int) error {
	if quantity < 1 {
		return fmt.Errorf("wrong quantity: %v", quantity)
	}
	d.orderDetails[item] = quantity
	fmt.Println(d.orderDetails)
	return nil
}
func (d *TakeAwayOrder) RemoveItem(item string) error {
	if _, ok := d.orderDetails[item]; !ok {
		return fmt.Errorf("wrong item: %v", item)
	}
	delete(d.orderDetails, item)
	fmt.Println(d.orderDetails)
	return nil
}
func (d *TakeAwayOrder) GetOrderDetails() map[string]int {
	return d.orderDetails
}

func ManageOrder(o Order) {
	o.AddItem("Pizza", 2)
	o.AddItem("Burger", 1)
	o.RemoveItem("Pizza")
	fmt.Println(o.GetOrderDetails())
}

func main() {
	dineIn := &DineInOrder{orderDetails: make(map[string]int)}
	takeAway := &TakeAwayOrder{orderDetails: make(map[string]int)}

	ManageOrder(dineIn)
	ManageOrder(takeAway)
}
