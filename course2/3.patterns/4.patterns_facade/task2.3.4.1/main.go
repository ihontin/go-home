package main

import (
	"errors"
	"fmt"
	"studentgit.kata.academy/Alkolex/go-kata/course2/2.oop/5.oop_mock/task2.2.5.1/pkg/indicators"
	"studentgit.kata.academy/Alkolex/go-kata/course2/2.oop/5.oop_mock/task2.2.5.1/pkg/models"
	"time"
)

// ---------------------------------------Indicatorer

////go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=Indicatorer
//
//type IndicatorWithCache struct {
//	indicator indicators.Indicatorer
//	cache     map[string][]float64
//}
//
//func (i *IndicatorWithCache) SMA(pair string, limit, period int, from, to time.Time) ([]float64, error) {
//	args := fmt.Sprintf("SMA, %s, %d, %d", pair, limit, period)
//	if sma, ok := i.cache[args]; ok {
//		return sma, nil
//	}
//	sma, err := i.indicator.SMA(pair, limit, period, from, to)
//	if err != nil {
//		return nil, err
//	}
//	i.cache[args] = sma
//	return sma, err
//}
//
//func (i *IndicatorWithCache) EMA(pair string, limit, period int, from, to time.Time) ([]float64, error) {
//	args := fmt.Sprintf("EMA, %s, %d, %d", pair, limit, period)
//
//	if sma, ok := i.cache[args]; ok {
//		return sma, nil
//	}
//	sma, err := i.indicator.EMA(pair, limit, period, from, to)
//	if err != nil {
//		return nil, err
//	}
//	i.cache[args] = sma
//	return sma, err
//}

//----------------------------------------------------Dashboarder

// Dashboarder должен возвращать историю свечей и индикаторы с несколькими периодами, заданными через opts

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=Dashboarder
type Dashboarder interface {
	GetDashboard(pair string, opts ...func(*Dashboard)) (DashboardData, error)
}

type DashboardData struct {
	Name           string
	CandlesHistory models.CandlesHistory
	Indicators     map[string][]IndicatorData
	limit          int
	from           time.Time
	to             time.Time
}

type IndicatorData struct {
	Name     string
	Period   int
	Indicate []float64
}

type IndicatorOpt struct {
	Name      string
	Periods   []int
	Indicator indicators.Indicatorer
}

type Dashboard struct {
	exchange           models.Exchanger
	withCandlesHistory bool
	IndicatorOpts      []IndicatorOpt
	limit              int
	from               time.Time
	to                 time.Time
}

func (d *Dashboard) GetDashboard(pair string, opts ...func(*Dashboard)) (DashboardData, error) {
	var newDashData = DashboardData{Indicators: make(map[string][]IndicatorData)}
	for _, val := range opts {
		val(d)
	}
	// Get CandlesHistory part
	if !d.withCandlesHistory {
		return DashboardData{}, errors.New("not enough options to find CandlesHistory")
	}
	ch, err := d.exchange.GetCandlesHistory(pair, d.limit, d.from, d.to)
	if err != nil {
		return DashboardData{}, err
	}
	newDashData.CandlesHistory = ch

	// Get Indicators map[string][]IndicatorData part
	for _, option := range d.IndicatorOpts {
		var indData = make([]IndicatorData, len(option.Periods))
		if option.Name == "SMA" {
			for i, period := range option.Periods {
				calced, err2 := option.Indicator.SMA(pair, d.limit, period, d.from, d.to)
				if err2 != nil {
					return DashboardData{}, err2
				}
				indData[i].Indicate = calced
				indData[i].Name = option.Name
				indData[i].Period = period
			}
		} else if option.Name == "EMA" {
			for i, period := range option.Periods {
				calced, err2 := option.Indicator.EMA(pair, d.limit, period, d.from, d.to)
				if err2 != nil {
					return DashboardData{}, err2
				}
				indData[i].Indicate = calced
				indData[i].Name = option.Name
				indData[i].Period = period
			}
		} else {
			return DashboardData{}, errors.New("option.Name not found")
		}
		newDashData.Indicators[option.Name] = indData
	}
	newDashData.Name = pair
	newDashData.limit = d.limit
	newDashData.from = d.from
	newDashData.to = d.to
	return newDashData, err
}

func WithCandlesHistory(limit int, from, to time.Time) func(*Dashboard) {
	return func(d *Dashboard) {
		d.limit = limit
		d.from = from
		d.to = to
		d.withCandlesHistory = true
	}
}

func WithIndicatorOpts(opts ...IndicatorOpt) func(*Dashboard) {
	return func(d *Dashboard) {
		for _, opt := range opts {
			d.IndicatorOpts = append(d.IndicatorOpts, opt)
		}
	}
}

func NewDashboard(exchange models.Exchanger) *Dashboard {
	return &Dashboard{exchange: exchange}
}

func main() {
	exchange := indicators.NewExmo()
	dashboard := NewDashboard(exchange)
	data, err := dashboard.GetDashboard("BTC_USD", WithCandlesHistory(30, time.Now().Add(-time.Hour*24*30), time.Now()), WithIndicatorOpts(
		IndicatorOpt{
			Name:      "SMA",
			Periods:   []int{5, 10, 20},
			Indicator: indicators.NewIndicator(exchange),
		},
		IndicatorOpt{
			Name:      "EMA",
			Periods:   []int{5, 10, 20},
			Indicator: indicators.NewIndicator(exchange),
		},
	))
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
