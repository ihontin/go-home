package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func GenProd(i int) []Product {
	var products []Product
	for j := 0; j < i; j++ {
		name := fmt.Sprintf("name-%v", j)
		price := float64(j)
		createAt := time.Now().Add(time.Hour * time.Duration(j))
		count := 200 + j
		products = append(products, Product{name, price, createAt, count})
	}
	return products
}

func TestByPrice_Len(t *testing.T) {
	a := []int{0, 3, 10}
	for i := range a {
		products := generateProducts(a[i])
		if ByPrice(products).Len() != a[i] {
			t.Errorf("expected = %v, got =  %v", a[i], ByPrice(products).Len())
		}
	}
}
func TestByPrice_Less(t *testing.T) {
	a := []int{3, 10}
	for _, i := range a {
		products := GenProd(i)
		if ByPrice(products).Less(0, 1) != true {
			t.Errorf("expected = %v, got =  %v", true, ByPrice(products).Less(0, 1))
		}
	}
}
func TestByPrice_Swap(t *testing.T) {
	a := []int{3, 10}
	for _, i := range a {
		products := GenProd(i)
		expected := products[0]
		ByPrice(products).Swap(0, 1)
		if !reflect.DeepEqual(expected, products[1]) {
			t.Errorf("expected = %v, got =  %v", expected, products[1])
		}
	}
}
func TestByCount_Len(t *testing.T) {
	a := []int{0, 3, 10}
	for i := range a {
		products := generateProducts(a[i])
		if ByCount(products).Len() != a[i] {
			t.Errorf("expected = %v, got =  %v", a[i], ByCount(products).Len())
		}
	}
}
func TestByCount_Less(t *testing.T) {
	a := []int{3, 10}
	for _, i := range a {
		products := GenProd(i)
		if ByCount(products).Less(0, 1) != true {
			t.Errorf("expected = %v, got =  %v", true, ByCount(products).Less(0, 1))
		}
	}
}
func TestByCount_Swap(t *testing.T) {
	a := []int{3, 10}
	for _, i := range a {
		products := GenProd(i)
		expected := products[0]
		ByCount(products).Swap(0, 1)
		if !reflect.DeepEqual(expected, products[1]) {
			t.Errorf("expected = %v, got =  %v", expected, products[1])
		}
	}
}

func TestByCreatedAt_Len(t *testing.T) {
	a := []int{0, 3, 10}
	for i := range a {
		products := generateProducts(a[i])
		if ByCreatedAt(products).Len() != a[i] {
			t.Errorf("expected = %v, got =  %v", a[i], ByCreatedAt(products).Len())
		}
	}
}
func TestByCreatedAt_Less(t *testing.T) {
	a := []int{3, 10}
	for _, i := range a {
		products := GenProd(i)
		if ByCreatedAt(products).Less(0, 1) != true {
			t.Errorf("expected = %v, got =  %v", true, ByCreatedAt(products).Less(0, 1))
		}
	}
}

func TestByCreatedAt_Swap(t *testing.T) {
	a := []int{3, 10}
	for _, i := range a {
		products := GenProd(i)
		expected := products[0]
		ByCreatedAt(products).Swap(0, 1)
		if !reflect.DeepEqual(expected, products[1]) {
			t.Errorf("expected = %v, got =  %v", expected, products[1])
		}
	}
}

func TestProduct_String(t *testing.T) {
	products := GenProd(20)
	got := products[0].String()
	expected := "Name: name-0, Price: 0.000000, Count: 200"
	if got != expected {
		t.Errorf("expected = %v, got =  %v", expected, got)
	}
}
