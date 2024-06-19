package main

import (
	"fmt"
	"net/http"
	"studentgit.kata.academy/Alkolex/go-kata/course2/2.oop/5.oop_mock/task2.2.5.1/pkg/models"

	//"studentgit.kata.academy/Alkolex/go-kata/course2/2.oop/5.oop_mock/task2.2.5.1/pkg/exmo"
	"studentgit.kata.academy/Alkolex/go-kata/course2/2.oop/5.oop_mock/task2.2.5.1/pkg/indicators"
	"time"
)

////go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=Indicatorer
//type Indicatorer interface {
//	SMA(pair string, limit, period int, from, to time.Time) ([]float64, error)
//	EMA(pair string, limit, period int, from, to time.Time) ([]float64, error)
//}

type IndicatorWithCache struct {
	indicator indicators.Indicatorer
	cache     map[string][]float64
}

func (i *IndicatorWithCache) SMA(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	args := fmt.Sprintf("SMA, %s, %d, %d", pair, limit, period)
	if sma, ok := i.cache[args]; ok {
		return sma, nil
	}
	sma, err := i.indicator.SMA(pair, limit, period, from, to)
	if err != nil {
		return nil, err
	}
	i.cache[args] = sma
	return sma, err
}

func (i *IndicatorWithCache) EMA(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	args := fmt.Sprintf("EMA, %s, %d, %d", pair, limit, period)

	if sma, ok := i.cache[args]; ok {
		return sma, nil
	}
	sma, err := i.indicator.EMA(pair, limit, period, from, to)
	if err != nil {
		return nil, err
	}
	i.cache[args] = sma
	return sma, err
}

func main() {
	var excha models.Exchanger
	urlEXMO := "https://api.exmo.com/v1.1"
	c := &http.Client{}
	excha = indicators.NewExmo(indicators.WithClient(c), indicators.WithURL(urlEXMO))
	var indicator = &IndicatorWithCache{
		indicator: indicators.NewIndicator(excha),
		cache:     make(map[string][]float64),
	}
	//time.Now().Add(-time.Hour*24)
	/*sma1*/
	sma1, _ := indicator.SMA("BTC_USD", 5, 5, time.Now().AddDate(0, 0, -2), time.Now())
	/*sma2*/ sma2, _ := indicator.SMA("BTC_USD", 5, 5, time.Now().AddDate(0, 0, -2), time.Now())
	/*ema1*/ ema1, _ := indicator.EMA("BTC_USD", 5, 5, time.Now().AddDate(0, 0, -2), time.Now())
	/*ema2*/ ema2, _ := indicator.EMA("BTC_USD", 5, 5, time.Now().AddDate(0, 0, -2), time.Now())
	fmt.Println(sma1[0], ema1[0])
	fmt.Println(sma2[0], ema2[0])
	//fmt.Print("29376.89 29376.89\n29376.89 29376.89\n")

}
