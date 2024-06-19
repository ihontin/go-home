package main

import (
	"reflect"
	"testing"
)

type testPrice struct {
	expected float64
	received []float64
}

func newProd() *Product {
	return &Product{
		ProductID:     1,
		Sells:         []float64{100, 200, 300},
		Buys:          []float64{50, 75, 100},
		CurrentPrice:  250,
		ProfitPercent: 10,
	}
}
func ThreProd() []*Product {
	var listProd = make([]*Product, 3)
	listProd[0] = &Product{
		ProductID:     0,
		Sells:         []float64{100, 200, 300},
		Buys:          []float64{50, 75, 100},
		CurrentPrice:  250,
		ProfitPercent: 10,
	}
	listProd[1] = &Product{
		ProductID:     1,
		Sells:         []float64{},
		Buys:          []float64{},
		CurrentPrice:  0,
		ProfitPercent: 0,
	}
	listProd[2] = &Product{
		ProductID:     2,
		Sells:         []float64{10, 20, 30},
		Buys:          []float64{10, 10, 10},
		CurrentPrice:  250,
		ProfitPercent: 10,
	}
	return listProd
}

func TestStatisticProfit_Sum(t *testing.T) {
	statistic := &StatisticProfit{}
	statistic.SetProduct(newProd())

	sumTest := []testPrice{
		{33.5, []float64{16, 2, 15.5}},
		{0, []float64{}},
		{0.2, []float64{0.2, 0}},
	}
	for _, val := range sumTest {
		if got := statistic.Sum(val.received); got != val.expected {
			t.Errorf("expected = %f, got = %f", val.expected, got)
		}
	}
}

func TestStatisticProfit_Average(t *testing.T) {
	statistic := &StatisticProfit{}
	statistic.SetProduct(newProd())
	sumTest := []testPrice{
		{11.166666666666666, []float64{16, 2, 15.5}},
		{0, []float64{}},
		{0.1, []float64{0.2, 0}},
	}
	for _, val := range sumTest {
		if got := statistic.Average(val.received); got != val.expected {
			t.Errorf("expected = %f, got = %f", val.expected, got)
		}
	}
}
func TestStatisticProfit_SetProduct(t *testing.T) {
	statistic := &StatisticProfit{}
	statistic.SetProduct(newProd())
	if got := statistic.product.ProductID; got != 1 {
		t.Errorf("expected = %d, got = %d", 1, got)
	}
}
func TestStatisticProfit_GetAverageProfit(t *testing.T) {
	statistic := &StatisticProfit{}
	expected := []float64{125, 0, 10.000000}
	forRange := ThreProd()
	for i, structTest := range forRange {
		statistic.SetProduct(structTest)
		if got := statistic.GetAverageProfit(); got != expected[i] {
			t.Errorf("expected = %f, got = %f", expected[i], got)
		}
	}
}

func TestStatisticProfit_GetAverageProfitPercent(t *testing.T) {
	statistic := &StatisticProfit{}
	expected := []float64{37.500000, 0, 50.000000}
	forRange := ThreProd()
	for i, structTest := range forRange {
		statistic.SetProduct(structTest)
		if got := statistic.GetAverageProfitPercent(); got != expected[i] {
			t.Errorf("expected = %f, got = %f", expected[i], got)
		}
	}
}

func TestStatisticProfit_GetCurrentProfit(t *testing.T) {
	statistic := &StatisticProfit{}
	expected := []float64{375.000000, 0, 30.000000}
	forRange := ThreProd()
	for i, structTest := range forRange {
		statistic.SetProduct(structTest)
		if got := statistic.GetCurrentProfit(); got != expected[i] {
			t.Errorf("expected = %f, got = %f", expected[i], got)
		}
	}
}

func TestStatisticProfit_GetDifferenceProfit(t *testing.T) {
	statistic := &StatisticProfit{}
	expected := []float64{250.000000, 0, 20.000000}
	forRange := ThreProd()
	for i, structTest := range forRange {
		statistic.SetProduct(structTest)
		if got := statistic.GetDifferenceProfit(); got != expected[i] {
			t.Errorf("expected = %f, got = %f", expected[i], got)
		}
	}
}

func TestStatisticProfit_GetAllData(t *testing.T) {
	statistic := &StatisticProfit{}
	expected := [][]float64{{125.000000, 37.500000, 375.000000, 250.000000},
		{0.000000, 0.000000, 0.000000, 0.000000}, {10.000000, 50.000000, 30.000000, 20.000000}}
	forRange := ThreProd()
	for i, structTest := range forRange {
		statistic.SetProduct(structTest)
		if got := statistic.GetAllData(); !reflect.DeepEqual(got, expected[i]) {
			t.Errorf("expected = %f, got = %f", expected[i], got)
		}
	}
}
