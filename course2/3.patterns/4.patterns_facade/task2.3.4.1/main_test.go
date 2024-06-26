package main

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"reflect"
	"studentgit.kata.academy/Alkolex/go-kata/course2/2.oop/5.oop_mock/task2.2.5.1/pkg/indicators"
	"testing"
	"time"
)

// Dashboarder is an autogenerated mock type for the Dashboarder type
type MockDashboarder struct {
	mock.Mock
}

// GetDashboard provides a mock function with given fields: pair, opts
func (_m *MockDashboarder) GetDashboard(pair string, opts ...func(*Dashboard)) (DashboardData, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pair)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 DashboardData
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...func(*Dashboard)) (DashboardData, error)); ok {
		return rf(pair, opts...)
	}
	if rf, ok := ret.Get(0).(func(string, ...func(*Dashboard)) DashboardData); ok {
		r0 = rf(pair, opts...)
	} else {
		r0 = ret.Get(0).(DashboardData)
	}

	if rf, ok := ret.Get(1).(func(string, ...func(*Dashboard)) error); ok {
		r1 = rf(pair, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDashboarder interface {
	mock.TestingT
	Cleanup(func())
}

// NewDashboarder creates a new instance of Dashboarder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDashboarder(t mockConstructorTestingTNewDashboarder) *MockDashboarder {
	mocke := &MockDashboarder{}
	mocke.Mock.Test(t)

	t.Cleanup(func() { mocke.AssertExpectations(t) })

	return mocke
}

func TestDashboard_GetDashboard(t *testing.T) {
	d := DashboardData{limit: 1, Name: "first"}
	tests := []struct {
		name    string
		want    DashboardData
		wantErr bool
		er      error
	}{
		{
			name:    "ok",
			want:    d,
			wantErr: false,
			er:      nil,
		},
		{
			name:    "!ok",
			want:    DashboardData{},
			wantErr: true,
			er:      errors.New("1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mDash := NewDashboarder(t)
			mDash.On("GetDashboard", "USD_BTC").Return(tt.want, tt.er)
			got, err := mDash.GetDashboard("USD_BTC")
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDashboard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDashboard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDashboard(t *testing.T) {
	exchange := indicators.NewExmo()
	dashboard := NewDashboard(exchange)
	if dashboard.exchange != exchange {
		t.Errorf("Expected = %v, got %v", exchange, dashboard.exchange)
	}
}

func TestWithCandlesHistory(t *testing.T) {
	exchange := indicators.NewExmo()
	dashboard := NewDashboard(exchange)
	WithCandlesHistory(30, time.Now().Add(-time.Hour*24*30), time.Now())(dashboard)
	if !dashboard.withCandlesHistory {
		t.Errorf("Expected = %v, got %v", true, dashboard.withCandlesHistory)
	}
}

func TestWithIndicatorOpts(t *testing.T) {
	exchange := indicators.NewExmo()
	dashboard := NewDashboard(exchange)
	WithIndicatorOpts(IndicatorOpt{
		Name:      "SMA",
		Periods:   []int{5, 10, 20},
		Indicator: indicators.NewIndicator(exchange),
	},
		IndicatorOpt{
			Name:      "EMA",
			Periods:   []int{5, 10, 20},
			Indicator: indicators.NewIndicator(exchange),
		})(dashboard)

	if len(dashboard.IndicatorOpts) != 2 {
		t.Errorf("Expected len = %d, got %d", 2, len(dashboard.IndicatorOpts))
	}
}
