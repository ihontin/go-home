package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	ticker         = "/ticker"
	trades         = "/trades"
	orderBook      = "/order_book"
	currency       = "/currency"
	candlesHistory = "/candles_history"
)

// Candle и CandlesHistory должен возвращать историю свечей по указанной паре валют.
type Candle struct {
	T int64   `json:"t"`
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V float64 `json:"v"`
}

type CandlesHistory struct {
	Candles []Candle `json:"candles"`
}

// Currencies должен возвращать список доступных валют.
type Currencies []string //валюты

// Pair и Trades должен возвращать список сделок по указанным парам валют.
type Pair struct {
	TradeID  int    `json:"trade_id"`
	Type     string `json:"type"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
	Amount   string `json:"amount"`
	Date     int    `json:"date"`
}
type Trades map[string][]Pair

// TickerValue и Ticker должен возвращать информацию о текущих курсах валют.
type TickerValue struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int    `json:"updated"`
}
type Ticker map[string]TickerValue

type OrderBookPair struct {
	AskQuantity string     `json:"ask_quantity"`
	AskAmount   string     `json:"ask_amount"`
	AskTop      string     `json:"ask_top"`
	BidQuantity string     `json:"bid_quantity"`
	BidAmount   string     `json:"bid_amount"`
	BidTop      string     `json:"bid_top"`
	Ask         [][]string `json:"ask"`
	Bid         [][]string `json:"bid"`
}
type Exchanger interface {
	GetTicker() (Ticker, error)
	GetTrades(pairs ...string) (Trades, error)
	GetOrderBook(limit int, pairs ...string) (OrderBook, error)
	GetCurrencies() (Currencies, error)
	GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error)
	GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error)
}

type OrderBook map[string]OrderBookPair

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

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error) {
	e.url = fmt.Sprintf("https://api.exmo.com/v1.1%s?symbol=%s&resolution=%d&from=%v&to=%v", candlesHistory, pair, limit, start.Unix(), end.Unix())
	resp, err := e.doRequest("GET", e.url, nil)
	if err != nil {
		fmt.Println("Ошибка отправке запроса:", err)
		return CandlesHistory{}, err
	}
	defer resp.Body.Close()

	var candHist = CandlesHistory{}
	err = json.NewDecoder(resp.Body).Decode(&candHist)
	if err != nil {
		fmt.Println("Ошибка декодирования запроса:", err)
		return CandlesHistory{}, err
	}
	return candHist, nil
}

func (e *Exmo) GetTicker() (Ticker, error) {
	resp, err := e.doRequest("POST", e.url+ticker, nil)
	if err != nil {
		fmt.Println("Ошибка отправке запроса:", err)
		return Ticker{}, err
	}
	defer resp.Body.Close()
	var ouTick = Ticker{}
	err = json.NewDecoder(resp.Body).Decode(&ouTick)
	if err != nil {
		fmt.Println("Ошибка декодирования запроса:", err)
		return Ticker{}, err
	}
	return ouTick, nil
}

func (e *Exmo) GetCurrencies() (Currencies, error) {
	resp, err := e.doRequest("POST", e.url+currency, nil)
	if err != nil {
		fmt.Println("Ошибка отправке запроса:", err)
		return Currencies{}, err
	}
	defer resp.Body.Close()
	var curOut = make(Currencies, 100)
	err = json.NewDecoder(resp.Body).Decode(&curOut)
	if err != nil {
		fmt.Println("Ошибка декодирования запроса:", err)
		return Currencies{}, err
	}
	return curOut, nil
}

func (e *Exmo) GetTrades(pairs ...string) (Trades, error) {
	structLen := len(pairs)
	resp, err := http.PostForm(e.url+trades, url.Values{"pair": pairs})
	if err != nil {
		fmt.Println("Ошибка отправке запроса:", err)
		return Trades{}, err
	}
	defer resp.Body.Close()
	var tradOut = make(Trades, structLen)
	err = json.NewDecoder(resp.Body).Decode(&tradOut)
	if err != nil {
		//fmt.Println("Ошибка декодирования запроса:", err)
		return Trades{}, err
	}
	return tradOut, nil
}

func (e *Exmo) GetOrderBook(limit int, pairs ...string) (OrderBook, error) {
	l := []string{fmt.Sprintf("%d", limit)}
	structLen := len(pairs)
	resp, err := http.PostForm(e.url+orderBook, url.Values{"pair": pairs, "limit": l})
	if err != nil {
		fmt.Println("Ошибка отправке запроса:", err)
		return OrderBook{}, err
	}
	defer resp.Body.Close()
	var ordBookOut = make(OrderBook, structLen)
	err = json.NewDecoder(resp.Body).Decode(&ordBookOut)
	if err != nil {
		//fmt.Println("Ошибка декодирования запроса:", err)
		return OrderBook{}, err
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

func main() {
	var exchange Exchanger
	exchange = NewExmo()
	/*periodCandles*/ _, err := exchange.GetCandlesHistory("BTC_USD", 30, time.Now().Add(-time.Hour*24), time.Now())
	if err != nil {
		return
	}
	//for _, c := range periodCandles.Candles {
	//	fmt.Println(c)
	//}
	urlEXMO := "https://api.exmo.com/v1.1"
	c := &http.Client{}
	exmoAll := NewExmo(WithClient(c), WithURL(urlEXMO))
	/*tick*/ _, err = exmoAll.GetTicker()
	if err != nil {
		log.Fatal("GetTicker error:", err)
		return
	}
	//fmt.Println(tick)
	/*cure*/
	_, err = exmoAll.GetCurrencies()
	if err != nil {
		log.Fatal("GetCurrencies error:", err)
		return
	}
	//fmt.Println(cure)
	/*trad*/
	_, err = exmoAll.GetTrades("BTC_USD,ETC_USD")
	if err != nil {
		log.Fatal("GetTrades error:", err)
		return
	}
	//for kay, val := range trad {
	//	fmt.Println(kay)
	//	for _, v := range val {
	//		fmt.Println(v)
	//	}
	//}
	/*orBook*/
	_, err = exmoAll.GetOrderBook(3, "BTC_USD,ETC_USD")
	if err != nil {
		log.Fatal("GetOrderBook error:", err)
		return
	}
	//for key, val := range orBook {
	//	fmt.Println(key, "ask: ", val.Ask)
	//}
	closePrice, err := exmoAll.GetClosePrice("ETC_USD", 30, time.Now().Add(-time.Hour*24), time.Now())
	if err != nil {
		log.Fatal("GetClosePrice error:", err)
		return
	}
	fmt.Println(closePrice)
}
