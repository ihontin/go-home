// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	models "studentgit.kata.academy/Alkolex/go-kata/course2/3.patterns/2.patterns_facade/task2.3.2.2/models"
)

// Githuber is an autogenerated mock type for the Githuber type
type Githuber struct {
	mock.Mock
}

// GetGists provides a mock function with given fields: ctx, username
func (_m *Githuber) GetGists(ctx context.Context, username string) ([]models.Item, error) {
	ret := _m.Called(ctx, username)

	var r0 []models.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]models.Item, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []models.Item); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Item)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRepos provides a mock function with given fields: ctx, username
func (_m *Githuber) GetRepos(ctx context.Context, username string) ([]models.Item, error) {
	ret := _m.Called(ctx, username)

	var r0 []models.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]models.Item, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []models.Item); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Item)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGithuber interface {
	mock.TestingT
	Cleanup(func())
}

// NewGithuber creates a new instance of Githuber. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGithuber(t mockConstructorTestingTNewGithuber) *Githuber {
	mock := &Githuber{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
