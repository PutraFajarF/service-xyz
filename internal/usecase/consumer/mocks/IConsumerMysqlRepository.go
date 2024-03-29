// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	entity "service-xyz/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// IConsumerMysqlRepository is an autogenerated mock type for the IConsumerMysqlRepository type
type IConsumerMysqlRepository struct {
	mock.Mock
}

// GetConsumerById provides a mock function with given fields: id
func (_m *IConsumerMysqlRepository) GetConsumerById(id int) (*entity.ConsumerInfo, error) {
	ret := _m.Called(id)

	var r0 *entity.ConsumerInfo
	if rf, ok := ret.Get(0).(func(int) *entity.ConsumerInfo); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.ConsumerInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertConsumer provides a mock function with given fields: data
func (_m *IConsumerMysqlRepository) InsertConsumer(data *entity.ConsumerInfo) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.ConsumerInfo) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIConsumerMysqlRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIConsumerMysqlRepository creates a new instance of IConsumerMysqlRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIConsumerMysqlRepository(t mockConstructorTestingTNewIConsumerMysqlRepository) *IConsumerMysqlRepository {
	mock := &IConsumerMysqlRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
