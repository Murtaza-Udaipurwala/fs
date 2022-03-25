// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IRepo is an autogenerated mock type for the IRepo type
type IRepo struct {
	mock.Mock
}

// Get provides a mock function with given fields: key
func (_m *IRepo) Get(key string) ([]byte, error) {
	ret := _m.Called(key)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: key, val
func (_m *IRepo) Set(key string, val []byte) error {
	ret := _m.Called(key, val)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte) error); ok {
		r0 = rf(key, val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
