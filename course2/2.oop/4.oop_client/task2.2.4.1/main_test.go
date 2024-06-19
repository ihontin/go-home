package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"
)

//go test -coverprofile=coverage.out
// go tool cover -func=coverage.out

func TestNewExmo(t *testing.T) {
	c := &http.Client{}
	urlExmo := "https://api.exmo.com/v1.1"
	exmoTest := NewExmo(WithClient(c), WithURL(urlExmo))
	if exmoTest.client != c {
		t.Errorf("expected = %v, got = %v", c, exmoTest.client)
	} else if exmoTest.url != urlExmo {
		t.Errorf("expected = %v, got = %v", urlExmo, exmoTest.url)
	}
}
func TestWithURL(t *testing.T) {
	urlExmo := "https://api.exmo.com/v1.1"
	exmoTest := &Exmo{}
	WithURL(urlExmo)(exmoTest)
	if exmoTest.url != urlExmo {
		t.Errorf("expected = %v, got = %v", urlExmo, exmoTest.url)
	}
}
func TestWithClient(t *testing.T) {
	c := &http.Client{}
	exmoTest := &Exmo{}
	WithClient(c)(exmoTest)
	if exmoTest.client != c {
		t.Errorf("expected = %v, got = %v", c, exmoTest.client)
	}
}
func TestExmo_doRequest(t *testing.T) {
	exmoTest := NewExmo(WithURL("https://api.exmo.com/v1.1"))
	_, err := exmoTest.doRequest("SET", "_-^", nil)
	if err == nil {
		t.Errorf("http.NewRequest error expected, but got nil")
	}
	a := io.Reader(os.Stdin)
	_, err = exmoTest.doRequest("DELETE", "https://v1.1/candles_history?symbol=BTC_USD&resolution=30&from=1585556979&to=1585557979", a)
	if err == nil {
		t.Errorf("client.Do(req) error expected, but got nil")
	}
	testExmo, _ := exmoTest.doRequest("GET", exmoTest.url+ticker, nil)
	defer testExmo.Body.Close()
	var ouTick = Ticker{}
	err = json.NewDecoder(testExmo.Body).Decode(&ouTick)
	if err != nil {
		t.Errorf("Decoding error")
	}
}
func TestExmo_GetCandlesHistory(t *testing.T) {
	e := NewExmo()
	structTest, err := e.GetCandlesHistory("BTC_USD", 30, time.Now().Add(-time.Hour*24), time.Now())
	if err == nil {
		fmt.Println("Error test")
		//t.Errorf("client.Do(req) error expected, but got nil")
	}
	var (
		i int64
		f float64
	)
	if reflect.TypeOf(structTest.Candles[0].T) != reflect.TypeOf(i) || reflect.TypeOf(structTest.Candles[0].O) != reflect.TypeOf(f) {
		t.Errorf("expected = %T, %T; got = %T, %T", i, f, structTest.Candles[0].T, structTest.Candles[0].O)
	}
}

func TestExmo_GetTicker(t *testing.T) {
	e := NewExmo(WithURL("https://api.exmo.com/v1.1"), WithClient(&http.Client{}))
	ticTest, err := e.GetTicker()
	if err == nil {
		fmt.Println("Error test")
		//t.Errorf("Decoding error expected, but got nil")
	}
	var (
		u int
		b string
	)
	for _, tt := range ticTest {
		if reflect.TypeOf(tt.Updated) != reflect.TypeOf(u) || reflect.TypeOf(tt.BuyPrice) != reflect.TypeOf(b) {
			t.Errorf("expected = %T, %T; got = %T, %T", u, b, tt.Updated, tt.BuyPrice)
		}
	}
}

func TestExmo_GetCurrencies(t *testing.T) {
	e := NewExmo(WithURL("https://api.exmo.com/v1.1"), WithClient(&http.Client{}))
	ticTest, err := e.GetCurrencies()
	if err == nil {
		fmt.Println("Error test")
		//t.Errorf("Decoding error expected, but got nil")
	}
	var testList = Currencies{"EXM", "USD", "EUR", "GBP", "PLN", "UAH", "USDT", "BTC", "LTC", "DOGE", "DASH", "ETH", "WAVES", "ZEC", "XRP", "ETC", "BCH", "EOS", "XLM", "OMG", "TRX", "ADA", "NEO", "GAS", "ZRX", "GUSD", "QTUM", "DAI", "MKR", "USDC", "ROOBEE", "XTZ", "VLX", "ONT", "ONG", "ALGO", "ATOM", "WXT", "CHZ", "ONE", "IQN", "PRQ", "HAI", "LINK", "UNI", "YFI", "GNY", "XYM", "DOT", "TON", "SGB", "SHIB", "GMT", "SOL", "EXFI", "SOLO", "NEAR", "DEBT", "PLCUC", "LYO", "ECS", "FLR", "PLCU", "SUI", "PEPE", "FLOKI", "EURT", "XAUT", "KAS"}

	for i, val := range testList {
		if val != ticTest[i] {
			t.Errorf("expected = %s; got = %s", val, ticTest[i])
		}
	}
}

func TestExmo_GetTrades(t *testing.T) {
	e := NewExmo(WithURL("https://api.exmo.com/v1.1"), WithClient(&http.Client{}))
	_, err := e.GetTrades()
	if err == nil {
		t.Errorf("Decoding error expected, but got nil")
	}
	tradTest, _ := e.GetTrades("BTC_USD,ETC_USD")
	var (
		tr int
		q  string
	)
	for _, val := range tradTest {
		for _, pair := range val {
			if reflect.TypeOf(pair.TradeID) != reflect.TypeOf(tr) || reflect.TypeOf(pair.Quantity) != reflect.TypeOf(q) {
				t.Errorf("expected = %T, %T; got = %T, %T", tr, q, pair.TradeID, pair.Quantity)
			}
		}
	}
}
func TestExmo_GetOrderBook(t *testing.T) {
	e := NewExmo(WithURL("https://api.exmo.com/v1.1"), WithClient(&http.Client{}))
	_, err := e.GetOrderBook(3)
	if err == nil {
		t.Errorf("Decoding error expected, but got nil")
	}
	tradTest, _ := e.GetOrderBook(3, "BTC_USD,ETC_USD")
	var (
		a  [][]string
		bq string
	)
	for _, pair := range tradTest {
		if reflect.TypeOf(pair.Ask) != reflect.TypeOf(a) || reflect.TypeOf(pair.BidQuantity) != reflect.TypeOf(bq) {
			t.Errorf("expected = %T, %T; got = %T, %T", a, bq, pair.Ask, pair.BidQuantity)
		}
	}
}

func TestExmo_GetClosePrice(t *testing.T) {
	e := NewExmo()
	structTest, err := e.GetClosePrice("BTC_USD", 30, time.Now().Add(-time.Hour*24), time.Now())
	if err == nil {
		fmt.Println("Error test")
		//t.Errorf("client.Do(req) error expected, but got nil")
	}
	var f float64
	if reflect.TypeOf(structTest[0]) != reflect.TypeOf(f) {
		t.Errorf("expected = %T; got = %T", f, structTest[0])
	}
}
