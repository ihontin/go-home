// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// GeneralIndicatorer is an autogenerated mock type for the GeneralIndicatorer type
type GeneralIndicatorer struct {
	mock.Mock
}

// GetData provides a mock function with given fields: pair, limit, period, from, to, indicator
func (_m *GeneralIndicatorer) GetData(pair string, limit int, period int, from time.Time, to time.Time, indicator Indicatorer) ([]float64, error) {
	ret := _m.Called(pair, limit, period, from, to, indicator)

	var r0 []float64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int, time.Time, time.Time, Indicatorer) ([]float64, error)); ok {
		return rf(pair, limit, period, from, to, indicator)
	}
	if rf, ok := ret.Get(0).(func(string, int, int, time.Time, time.Time, Indicatorer) []float64); ok {
		r0 = rf(pair, limit, period, from, to, indicator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]float64)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int, time.Time, time.Time, Indicatorer) error); ok {
		r1 = rf(pair, limit, period, from, to, indicator)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGeneralIndicatorer interface {
	mock.TestingT
	Cleanup(func())
}

// NewGeneralIndicatorer creates a new instance of GeneralIndicatorer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGeneralIndicatorer(t mockConstructorTestingTNewGeneralIndicatorer) *GeneralIndicatorer {
	mock := &GeneralIndicatorer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
