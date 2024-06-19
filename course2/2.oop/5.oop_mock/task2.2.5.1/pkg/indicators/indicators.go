package indicators

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"studentgit.kata.academy/Alkolex/go-kata/course2/2.oop/5.oop_mock/task2.2.5.1/pkg/models"
	"time"
)

const (
	ticker         = "/ticker"
	trades         = "/trades"
	orderBook      = "/order_book"
	currency       = "/currency"
	candlesHistory = "/candles_history"
)

// Exmo структура для работы с ресурсом Exmo
type Exmo struct {
	client *http.Client
	url    string
}

// NewExmo Конструктор структуры Exmo
func NewExmo(opts ...func(exmo *Exmo)) *Exmo {
	outExmo := &Exmo{client: &http.Client{}}
	for _, opt := range opts {
		opt(outExmo)
	}
	return outExmo
}

func WithClient(client *http.Client) func(exmo *Exmo) {
	return func(exmo *Exmo) {
		exmo.client = client
	}
}

func WithURL(url string) func(exmo *Exmo) {
	return func(exmo *Exmo) {
		exmo.url = url
	}
}

// отправляет запрос на сервер, возвращает ответ *http.Response
func (e *Exmo) doRequest(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	resp, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end time.Time) (models.CandlesHistory, error) {
	e.url = fmt.Sprintf("https://api.exmo.com/v1.1%s?symbol=%s&resolution=%d&from=%v&to=%v", candlesHistory, pair, limit, start.Unix(), end.Unix())

	resp, err := e.doRequest("GET", e.url, nil)
	if err != nil {
		fmt.Println("Ошибка отправке запроса:", err)
		return models.CandlesHistory{}, err
	}

	defer resp.Body.Close()

	var candHist = models.CandlesHistory{}
	err = json.NewDecoder(resp.Body).Decode(&candHist)
	if err != nil {
		fmt.Println("Ошибка декодирования запроса:", err)
		return models.CandlesHistory{}, err
	}
	return candHist, nil
}

func (e *Exmo) GetTicker() (models.Ticker, error) {
	resp, err := e.doRequest("POST", e.url+ticker, nil)
	if err != nil {
		fmt.Println("Ошибка отправке запроса:", err)
		return models.Ticker{}, err
	}
	defer resp.Body.Close()
	var ouTick = models.Ticker{}
	err = json.NewDecoder(resp.Body).Decode(&ouTick)
	if err != nil {
		fmt.Println("Ошибка декодирования запроса:", err)
		return models.Ticker{}, err
	}
	return ouTick, nil
}

func (e *Exmo) GetCurrencies() (models.Currencies, error) {
	resp, err := e.doRequest("POST", e.url+currency, nil)
	if err != nil {
		fmt.Println("Ошибка отправке запроса:", err)
		return models.Currencies{}, err
	}
	defer resp.Body.Close()
	var curOut = make(models.Currencies, 100)
	err = json.NewDecoder(resp.Body).Decode(&curOut)
	if err != nil {
		fmt.Println("Ошибка декодирования запроса:", err)
		return models.Currencies{}, err
	}
	return curOut, nil
}

func (e *Exmo) GetTrades(pairs ...string) (models.Trades, error) {
	structLen := len(pairs)
	resp, err := http.PostForm(e.url+trades, url.Values{"pair": pairs})
	if err != nil {
		fmt.Println("Ошибка отправке запроса:", err)
		return models.Trades{}, err
	}
	defer resp.Body.Close()
	var tradOut = make(models.Trades, structLen)
	err = json.NewDecoder(resp.Body).Decode(&tradOut)
	if err != nil {
		//fmt.Println("Ошибка декодирования запроса:", err)
		return models.Trades{}, err
	}
	return tradOut, nil
}

func (e *Exmo) GetOrderBook(limit int, pairs ...string) (models.OrderBook, error) {
	l := []string{fmt.Sprintf("%d", limit)}
	structLen := len(pairs)
	resp, err := http.PostForm(e.url+orderBook, url.Values{"pair": pairs, "limit": l})
	if err != nil {
		fmt.Println("Ошибка отправке запроса:", err)
		return models.OrderBook{}, err
	}
	defer resp.Body.Close()
	var ordBookOut = make(models.OrderBook, structLen)
	err = json.NewDecoder(resp.Body).Decode(&ordBookOut)
	if err != nil {
		//fmt.Println("Ошибка декодирования запроса:", err)
		return models.OrderBook{}, err
	}
	return ordBookOut, nil
}

func (e *Exmo) GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error) {
	cHistory, err := e.GetCandlesHistory(pair, limit, start, end)
	if err != nil {
		fmt.Println("Ошибка GetCandlesHistory:", err)
		return []float64{}, err
	}
	var closePrice = make([]float64, 0, len(cHistory.Candles))
	for _, candle := range cHistory.Candles {
		closePrice = append(closePrice, candle.C)

	}

	return closePrice, nil

}

type Indicatorer interface {
	SMA(pair string, limit, period int, from, to time.Time) ([]float64, error)
	EMA(pair string, limit, period int, from, to time.Time) ([]float64, error)
}

type Indicator struct {
	exchange     models.Exchanger
	calculateSMA func(data []float64, period int) []float64
	calculateEMA func(data []float64, period int) []float64
}

func (i *Indicator) SMA(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	closePrice, err := i.exchange.GetClosePrice(pair, limit, from, to)
	if err != nil {
		return []float64{}, err
	}
	return calculateSMA(closePrice, period), err
}
func (i *Indicator) EMA(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	closePrice, err := i.exchange.GetClosePrice(pair, limit, from, to)
	if err != nil {
		return nil, err
	}
	return calculateEMA(closePrice, period), err
}

func calculateSMA(data []float64, window int) []float64 {
	if len(data) == 0 || window == 0 {
		return []float64{}
	}
	sma := make([]float64, len(data))
	windowSum := 0.0

	for i := 0; i < len(data); i++ {
		windowSum += data[i]
		if i >= window {
			windowSum -= data[i-window]
		}
		sma[i] = windowSum / float64(minAB(i+1, window))
	}

	return sma
}

// Вспомогательная функция для нахождения минимума двух чисел
func minAB(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Функция для расчета EMA
func calculateEMA(data []float64, window int) []float64 {
	if len(data) == 0 || window == 0 {
		return []float64{}
	}
	ema := make([]float64, len(data))
	multiplier := 2.0 / float64(window+1)

	// Инициализация первого значения EMA
	ema[0] = data[0]

	for i := 1; i < len(data); i++ {
		ema[i] = (data[i]-ema[i-1])*multiplier + ema[i-1]
	}

	return ema
}

type IndicatorOption func(*Indicator)

func WithSMA(calculateSMA func(data []float64, period int) []float64) func(*Indicator) {
	return func(indicator *Indicator) {
		indicator.calculateSMA = calculateSMA
	}
}

func WithEMA(calculateEMA func(data []float64, period int) []float64) func(*Indicator) {
	return func(indicator *Indicator) {
		indicator.calculateEMA = calculateEMA
	}
}

func NewIndicator(exchange models.Exchanger, opts ...IndicatorOption) *Indicator {
	i := &Indicator{
		exchange:     exchange,
		calculateEMA: calculateEMA,
		calculateSMA: calculateSMA,
	}
	for _, opt := range opts {
		opt(i)
	}
	return i
}
