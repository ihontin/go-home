// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

//import mock "github.com/stretchr/testify/mock"

// Dashboarder is an autogenerated mock type for the Dashboarder type
//type Dashboarder struct {
//	mock.Mock
//}
//
//// GetDashboard provides a mock function with given fields: pair, opts
//func (_m *Dashboarder) GetDashboard(pair string, opts ...func(*Dashboard)) (DashboardData, error) {
//	_va := make([]interface{}, len(opts))
//	for _i := range opts {
//		_va[_i] = opts[_i]
//	}
//	var _ca []interface{}
//	_ca = append(_ca, pair)
//	_ca = append(_ca, _va...)
//	ret := _m.Called(_ca...)
//
//	var r0 DashboardData
//	var r1 error
//	if rf, ok := ret.Get(0).(func(string, ...func(*Dashboard)) (DashboardData, error)); ok {
//		return rf(pair, opts...)
//	}
//	if rf, ok := ret.Get(0).(func(string, ...func(*Dashboard)) DashboardData); ok {
//		r0 = rf(pair, opts...)
//	} else {
//		r0 = ret.Get(0).(DashboardData)
//	}
//
//	if rf, ok := ret.Get(1).(func(string, ...func(*Dashboard)) error); ok {
//		r1 = rf(pair, opts...)
//	} else {
//		r1 = ret.Error(1)
//	}
//
//	return r0, r1
//}
//
//type mockConstructorTestingTNewDashboarder interface {
//	mock.TestingT
//	Cleanup(func())
//}
//
//// NewDashboarder creates a new instance of Dashboarder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
//func NewDashboarder(t mockConstructorTestingTNewDashboarder) *Dashboarder {
//	mock := &Dashboarder{}
//	mock.Mock.Test(t)
//
//	t.Cleanup(func() { mock.AssertExpectations(t) })
//
//	return mock
//}