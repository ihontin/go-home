package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"
	"studentgit.kata.academy/Alkolex/go-kata/course2/3.patterns/2.patterns_facade/task2.3.2.1/mocks"
	"testing"
	"time"
)

func TestFuncMain(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	expected := "29376.89 29376.89\n29376.89 29376.89\n"
	stdout := bytes.Buffer{}
	_, _ = stdout.ReadFrom(r)
	if expected != stdout.String() {
		t.Errorf("expected = %s, got = %s", expected, stdout.String())
	}
}

func TestIndicatorWithCache_SMA(t *testing.T) {
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
			name: "some test",
			args: args{pair: "BTC_USD",
				limit:  5,
				period: 5,
				from:   time.Now().AddDate(0, 0, -2),
				to:     time.Now(),
			},
			want:    []float64{821.2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			argsCache := fmt.Sprintf("EMA, %s, %d, %d", tt.args.pair, tt.args.limit, tt.args.period)
			indiEMA := mocks.NewIndicatorer(t)
			indiEMA.On("SMA", tt.args.pair, tt.args.limit, tt.args.period, tt.args.from, tt.args.to).
				Return([]float64{821.2}, nil)
			i := &IndicatorWithCache{
				indicator: indiEMA,
				cache:     make(map[string][]float64),
			}
			if sma, ok := i.cache[argsCache]; ok {
				t.Errorf("SMA() got = %v, want %v", sma, []float64{})
				return
			}

			got, err := i.SMA(tt.args.pair, tt.args.limit, tt.args.period, tt.args.from, tt.args.to)
			i.cache[argsCache] = got
			if (err == nil) == tt.wantErr {
				t.Errorf("SMA() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SMA() got = %v, want %v", got, tt.want)
			}
			if sma, ok := i.cache[argsCache]; !ok {
				t.Errorf("SMA() got = %v, want %v", sma, []float64{821.2})
				return
			}

			indiEMA2 := mocks.NewIndicatorer(t)
			indiEMA2.On("SMA", tt.args.pair, tt.args.limit, tt.args.period, tt.args.from, tt.args.to).
				Return([]float64{}, errors.New("SMA error"))
			ii := &IndicatorWithCache{
				indicator: indiEMA2,
				cache:     make(map[string][]float64),
			}
			got, err = ii.SMA(tt.args.pair, tt.args.limit, tt.args.period, tt.args.from, tt.args.to)
			if (err == nil) == true {
				t.Errorf("SMA() error = %v, wantErr %v", err, true)
				return
			}
		})
	}
}

func TestIndicatorWithCache_EMA(t *testing.T) {
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
			name: "some test",
			args: args{pair: "BTC_USD",
				limit:  5,
				period: 5,
				from:   time.Now().AddDate(0, 0, -2),
				to:     time.Now(),
			},
			want:    []float64{821.2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			argsCache := fmt.Sprintf("EMA, %s, %d, %d", tt.args.pair, tt.args.limit, tt.args.period)
			indiEMA := mocks.NewIndicatorer(t)
			indiEMA.On("EMA", tt.args.pair, tt.args.limit, tt.args.period, tt.args.from, tt.args.to).
				Return([]float64{821.2}, nil)
			i := &IndicatorWithCache{
				indicator: indiEMA,
				cache:     make(map[string][]float64),
			}
			if ema, ok := i.cache[argsCache]; ok {
				t.Errorf("EMA() got = %v, want %v", ema, []float64{})
				return
			}

			got, err := i.EMA(tt.args.pair, tt.args.limit, tt.args.period, tt.args.from, tt.args.to)
			i.cache[argsCache] = got
			if (err == nil) == tt.wantErr {
				t.Errorf("EMA() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EMA() got = %v, want %v", got, tt.want)
			}
			if ema, ok := i.cache[argsCache]; !ok {
				t.Errorf("EMA() got = %v, want %v", ema, []float64{821.2})
				return
			}

			indiEMA2 := mocks.NewIndicatorer(t)
			indiEMA2.On("EMA", tt.args.pair, tt.args.limit, tt.args.period, tt.args.from, tt.args.to).
				Return([]float64{}, errors.New("EMA error"))
			ii := &IndicatorWithCache{
				indicator: indiEMA2,
				cache:     make(map[string][]float64),
			}
			got, err = ii.EMA(tt.args.pair, tt.args.limit, tt.args.period, tt.args.from, tt.args.to)
			if (err == nil) == true {
				t.Errorf("EMA() error = %v, wantErr %v", err, true)
				return
			}
		})
	}
}
