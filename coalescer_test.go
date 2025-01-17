// Code generated by mockery v2.12.0. DO NOT EDIT.

package goalesce

import (
	reflect "reflect"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// mockCoalescer is an autogenerated mock type for the Coalescer type
type mockCoalescer struct {
	mock.Mock
}

// Coalesce provides a mock function with given fields: v1, v2
func (_m *mockCoalescer) Coalesce(v1 reflect.Value, v2 reflect.Value) (reflect.Value, error) {
	ret := _m.Called(v1, v2)

	var r0 reflect.Value
	if rf, ok := ret.Get(0).(func(reflect.Value, reflect.Value) reflect.Value); ok {
		r0 = rf(v1, v2)
	} else {
		r0 = ret.Get(0).(reflect.Value)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(reflect.Value, reflect.Value) error); ok {
		r1 = rf(v1, v2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithFallback provides a mock function with given fields: fallback
func (_m *mockCoalescer) WithFallback(fallback Coalescer) {
	_m.Called(fallback)
}

// newMockCoalescer creates a new instance of mockCoalescer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func newMockCoalescer(t testing.TB) *mockCoalescer {
	mock := &mockCoalescer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
