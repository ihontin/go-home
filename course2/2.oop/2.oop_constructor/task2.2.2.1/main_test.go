package main

import (
	"reflect"
	"testing"
	"time"
)

func TestWithOrderDate(t *testing.T) {
	expecteds := []time.Time{time.Now(), time.Time{}}
	testOrder := &Order{}
	for _, expected := range expecteds {
		WithOrderDate(expected)(testOrder)
		if expected != testOrder.OrderDate {
			t.Errorf("expected = %v, got = %v", expected, testOrder.OrderDate)
		}
	}
}

func TestWithItems(t *testing.T) {
	expecteds := [][]string{[]string{"item1", "item2"}, []string{}}
	testOrder := &Order{}
	for _, expected := range expecteds {
		WithItems(expected)(testOrder)
		if !reflect.DeepEqual(expected, testOrder.Items) {
			t.Errorf("expected = %v, got = %v", expected, testOrder.Items)
		}
	}
}

func TestWithCustomerID(t *testing.T) {
	expecteds := []string{"item1", "item2", ""}
	testOrder := &Order{}
	for _, expected := range expecteds {
		WithCustomerID(expected)(testOrder)
		if expected != testOrder.CustomerID {
			t.Errorf("expected = %v, got = %v", expected, testOrder.CustomerID)
		}
	}
}

func TestNewOrder(t *testing.T) {
	idTest := 21
	tiTest := time.Now()
	custIdTest := "21"
	itemTest := []string{"item1", "item2"}
	expected := NewOrder(idTest,
		WithCustomerID(custIdTest),
		WithItems(itemTest),
		WithOrderDate(tiTest))
	testO := &Order{idTest, custIdTest, itemTest, tiTest}
	if expected.ID != testO.ID {
		t.Errorf("expected = %v, got = %v", expected.ID, testO.ID)
	}
	if expected.OrderDate != testO.OrderDate {
		t.Errorf("expected = %v, got = %v", expected.OrderDate, testO.OrderDate)
	}
	if expected.CustomerID != testO.CustomerID {
		t.Errorf("expected = %v, got = %v", expected.CustomerID, testO.CustomerID)
	}
	if !reflect.DeepEqual(expected.Items, testO.Items) {
		t.Errorf("expected = %v, got = %v", expected.Items, testO.Items)
	}
}
