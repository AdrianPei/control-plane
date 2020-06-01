// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// KubeconfigProvider is an autogenerated mock type for the KubeconfigProvider type
type KubeconfigProvider struct {
	mock.Mock
}

// FetchRaw provides a mock function with given fields: shootName
func (_m *KubeconfigProvider) FetchRaw(shootName string) ([]byte, error) {
	ret := _m.Called(shootName)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(shootName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(shootName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}