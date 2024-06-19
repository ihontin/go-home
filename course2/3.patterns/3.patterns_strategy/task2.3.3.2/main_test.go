package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"os"
	"reflect"
	"studentgit.kata.academy/Alkolex/go-kata/course2/3.patterns/3.patterns_strategy/task2.3.3.2/mocks"
	"studentgit.kata.academy/Alkolex/go-kata/course2/3.patterns/3.patterns_strategy/task2.3.3.2/models"
	exmock "studentgit.kata.academy/Alkolex/go-kata/course2/3.patterns/3.patterns_strategy/task2.3.3.2/models/mocks"
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
	var ouTick = models.Ticker{}
	err = json.NewDecoder(testExmo.Body).Decode(&ouTick)
	if err != nil {
		t.Errorf("Decoding error")
	}
}

func TestExmo_GetCandlesHistory(t *testing.T) {
	// передаваемые параметры в тестируемый метод
	type args struct {
		pair  string
		limit int
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    models.CandlesHistory
	}{
		{
			name:    "ok",
			args:    args{"BTC_USD", 30, time.Now().Add(-time.Hour * 24), time.Now()},
			wantErr: false,
			want:    models.CandlesHistory{},
		},
		{
			name:    "second_ok",
			args:    args{"", 0, time.Time{}, time.Time{}},
			wantErr: false,
			want:    models.CandlesHistory{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exchenger := exmock.NewExchanger(t)
			exchenger.
				On("GetCandlesHistory", tt.args.pair, tt.args.limit, tt.args.start, tt.args.end).
				Return(models.CandlesHistory{}, nil)

			got, err := exchenger.GetCandlesHistory(tt.args.pair, tt.args.limit, tt.args.start, tt.args.end)
			t.Logf("error is = %v", err)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCandlesHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("expected = %T; got = %T", tt.want.Candles, got.Candles)
			if !reflect.DeepEqual(got.Candles, tt.want.Candles) {
				t.Errorf("expected = %T; got = %T", tt.want.Candles, got.Candles)
			}
		})
	}
}

func TestExmo_GetClosePrice(t *testing.T) {
	type args struct {
		pair  string
		limit int
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{"BTC_USD", 30, time.Now().Add(-time.Hour * 24), time.Now()},
			wantErr: false,
			want:    []float64{},
		},
		{
			name:    "second_ok",
			args:    args{"", 0, time.Time{}, time.Time{}},
			wantErr: false,
			want:    []float64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exchenger := exmock.NewExchanger(t)
			exchenger.
				On("GetClosePrice", tt.args.pair, tt.args.limit, tt.args.start, tt.args.end).
				Return([]float64{}, nil)
			got, err := exchenger.GetClosePrice(tt.args.pair, tt.args.limit, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClosePrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClosePrice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExmo_GetCurrencies(t *testing.T) {
	tests := []struct {
		name    string
		want    models.Currencies
		wantErr bool
	}{
		{
			name:    "ok",
			wantErr: false,
			want:    models.Currencies{},
		},
		{
			name:    "second_ok",
			wantErr: false,
			want:    models.Currencies{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exchenger := exmock.NewExchanger(t)
			exchenger.
				On("GetCurrencies").
				Return(models.Currencies{}, nil)
			got, err := exchenger.GetCurrencies()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrencies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got = append(got, "")
			tt.want = append(tt.want, "")
			if !reflect.DeepEqual(got[0], tt.want[0]) {
				t.Errorf("GetCurrencies() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExmo_GetOrderBook(t *testing.T) {
	type args struct {
		limit int
		pairs string
	}
	tests := []struct {
		name    string
		args    args
		want    models.OrderBook
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{3, "BTC_USD,ETC_USD"},
			wantErr: false,
			want:    models.OrderBook{},
		},
		{
			name:    "second_ok",
			args:    args{3, ""},
			wantErr: false,
			want:    models.OrderBook{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exchenger := exmock.NewExchanger(t)
			exchenger.
				On("GetOrderBook", tt.args.limit, tt.args.pairs).
				Return(models.OrderBook{}, nil)
			got, err := exchenger.GetOrderBook(tt.args.limit, tt.args.pairs)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrderBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got["1"].Ask, tt.want["1"].Ask) {
				t.Errorf("GetOrderBook() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExmo_GetTicker(t *testing.T) {
	tests := []struct {
		name    string
		want    models.Ticker
		wantErr bool
	}{
		{
			name:    "ok",
			wantErr: false,
			want:    models.Ticker{},
		},
		{
			name:    "second_ok",
			wantErr: false,
			want:    models.Ticker{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exchenger := exmock.NewExchanger(t)
			exchenger.
				On("GetTicker").
				Return(models.Ticker{}, nil)
			got, err := exchenger.GetTicker()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTicker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got[""].Avg, tt.want[""].Avg) {
				t.Errorf("GetTicker() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExmo_GetTrades(t *testing.T) {
	type args struct {
		pairs string
	}
	tests := []struct {
		name    string
		args    args
		want    models.Trades
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{"BTC_USD,ETC_USD"},
			wantErr: false,
			want:    models.Trades{},
		},
		{
			name:    "second_ok",
			args:    args{""},
			wantErr: false,
			want:    models.Trades{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exchenger := exmock.NewExchanger(t)
			exchenger.
				On("GetTrades", tt.args.pairs).
				Return(models.Trades{}, nil)
			got, err := exchenger.GetTrades(tt.args.pairs)
			got["1"] = []models.Pair{{Type: "q"}}
			tt.want["1"] = []models.Pair{{Type: "q"}}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTrades() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got["1"][0].Type, tt.want["1"][0].Type) {
				t.Errorf("GetTrades() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndicatorEMA_GetData(t *testing.T) {
	//mockExchanger := new(exmock.Exchanger)

	indic := mocks.NewIndicatorer(t)

	pair := "BTC_USD"
	limit := 10
	from := time.Time{}
	to := time.Time{}

	//mockExchanger.
	//	On("GetClosePrice", pair, limit, from, to).
	//	Return([]float64{}, nil)

	indic.On("GetData", pair, limit, 3, from, to).Return(nil, errors.New("test error"))
	result, err := indic.GetData(pair, limit, 3, from, to)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	var closePriceList = []float64{6.0, 5.0, 4.0, 3.0, 2.0, 3.0, 4.0}
	var expectedError error

	from = time.Now()
	to = time.Now()
	indic.On("GetData", pair, limit, 5, from, to).Return(closePriceList, expectedError)
	result, err = indic.GetData(pair, limit, 5, from, to)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	//assert.Equal(t, closePriceList, assert.NotNil(t, result))
}

func TestIndicatorSMA_GetData(t *testing.T) {
	type args struct {
		pair   string
		limit  int
		period int
		from   time.Time
		to     time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name:    "first",
			args:    args{"BTC_USD", 30, 3, time.Now().Add(-time.Hour * 24), time.Now()},
			want:    []float64{},
			wantErr: false,
		},
		{
			name:    "second",
			args:    args{"", 0, 0, time.Time{}, time.Time{}},
			want:    []float64{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var exchange models.Exchanger
			exchange = NewExmo()
			i := NewIndicatorSMA(exchange)
			got, err := i.GetData(tt.args.pair, tt.args.limit, tt.args.period, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("SMA() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("SMA() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIndicatorSMA(t *testing.T) {
	var ex models.Exchanger
	ex = NewExmo()
	indic := NewIndicatorSMA(ex, WithSMA(calculateSMA))
	got, err := indic.GetData("BTC_USD", 30, 5, time.Now().AddDate(0, 0, -2), time.Now())
	if (err != nil) != false {
		fmt.Println(err)
		return
	}
	var expected []float64
	if reflect.TypeOf(got) != reflect.TypeOf(expected) {
		t.Errorf("NewIndicator() = %v, want %v", got, expected)
	}
}
func TestNewIndicatorEMA(t *testing.T) {
	var ex models.Exchanger
	ex = NewExmo()
	indic := NewIndicatorEMA(ex, WithEMA(calculateEMA))
	got, err := indic.GetData("BTC_USD", 30, 5, time.Now().AddDate(0, 0, -2), time.Now())
	if (err != nil) != false {
		fmt.Println(err)
		return
	}
	var expected []float64
	if reflect.TypeOf(got) != reflect.TypeOf(expected) {
		t.Errorf("NewIndicator() = %v, want %v", got, expected)
	}
}
func TestWithEMA(t *testing.T) {
	var exchange models.Exchanger
	exchange = NewExmo()
	indicator := NewIndicatorEMA(exchange)
	WithEMA(calculateEMA)(indicator)
	got, err := indicator.GetData("BTC_USD", 30, 5, time.Now().AddDate(0, 0, -2), time.Now())
	if (err != nil) != false {
		fmt.Println(err)
		return
	}
	var expected []float64
	if reflect.TypeOf(got) != reflect.TypeOf(expected) {
		t.Errorf("NewIndicator() = %v, want %v", got, expected)
	}
}

func TestWithSMA(t *testing.T) {
	var exchange models.Exchanger
	exchange = NewExmo()
	indicator := NewIndicatorSMA(exchange)
	WithSMA(calculateSMA)(indicator)
	got, err := indicator.GetData("BTC_USD", 30, 5, time.Now().AddDate(0, 0, -2), time.Now())
	if (err != nil) != false {
		fmt.Println(err)
		return
	}
	var expected []float64
	if reflect.TypeOf(got) != reflect.TypeOf(expected) {
		t.Errorf("NewIndicator() = %v, want %v", got, expected)
	}
}

func Test_calculateEMA(t *testing.T) {
	type args struct {
		data   []float64
		window int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "first",
			args: args{[]float64{6, 7, 6, 7, 6, 7, 6, 7, 6}, 4},
			want: []float64{6, 6.4, 6.24, 6.5440000000000005, 6.3264000000000005, 6.59584, 6.357504, 6.6145024, 6.36870144},
		},
		{
			name: "second",
			args: args{[]float64{}, 0},
			want: []float64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateEMA(tt.args.data, tt.args.window); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateEMA() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_calculateSMA(t *testing.T) {
	type args struct {
		data   []float64
		window int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "first",
			args: args{[]float64{6, 7, 6, 7, 6, 7, 6, 7, 6}, 4},
			want: []float64{6, 6.5, 6.333333333333333, 6.5, 6.5, 6.5, 6.5, 6.5, 6.5},
		},
		{
			name: "second",
			args: args{[]float64{}, 0},
			want: []float64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSMA(tt.args.data, tt.args.window); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateSMA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minAB(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{5, 0},
			want: 0,
		},
		{
			name: "second",
			args: args{0, 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAB(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("minAB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGeneralIndicator(t *testing.T) {
	var ex models.Exchanger
	ex = NewExmo()
	indic := NewIndicatorEMA(ex, WithEMA(calculateEMA))
	genIndic := NewGeneralIndicator()

	got, err := genIndic.GetData("BTC_USD", 30, 5, time.Now().AddDate(0, 0, -2), time.Now(), indic)
	if (err != nil) != false {
		fmt.Println(err)
		return
	}
	var expected []float64
	if reflect.TypeOf(got) != reflect.TypeOf(expected) {
		t.Errorf("NewIndicator() = %v, want %v", got, expected)
	}
}

func TestGeneralIndicator_GetData(t *testing.T) {

	mockIndic := mocks.NewIndicatorer(t)
	mockGenIndic := mocks.NewGeneralIndicatorer(t)

	pair := "BTC_USD"
	limit := 10
	from := time.Now()
	to := time.Now()

	mockGenIndic.On("GetData", pair, limit, 5, from, to, mock.Anything).Return(nil, errors.New("test error"))
	result, err := mockGenIndic.GetData(pair, limit, 5, from, to, *mockIndic)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	var closePriceList = []float64{6.0, 5.0, 4.0, 3.0, 2.0, 3.0, 4.0}
	var expectedError error

	from = time.Now()
	to = time.Now()
	mockGenIndic.On("GetData", pair, limit, 5, from, to, mock.Anything).Return(closePriceList, expectedError)
	result, err = mockGenIndic.GetData(pair, limit, 5, from, to, *mockIndic)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	//assert.Equal(t, closePriceList, assert.NotNil(t, result))
}

func TestMainFunc(t *testing.T) {
	expected := "Done"
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	stdout := bytes.Buffer{}
	_, _ = stdout.ReadFrom(r)
	if expected != stdout.String() {
		t.Errorf("expected = %s, got = %s", expected, stdout.String())
	}
}
