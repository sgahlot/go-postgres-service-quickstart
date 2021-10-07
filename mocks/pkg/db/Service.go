// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	db "github.com/sgahlot/go-postgres-service-quickstart/pkg/db"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// DeleteFruits provides a mock function with given fields: req
func (_m *Service) DeleteFruits(req *db.FruitRequest) db.FruitResponse {
	ret := _m.Called(req)

	var r0 db.FruitResponse
	if rf, ok := ret.Get(0).(func(*db.FruitRequest) db.FruitResponse); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(db.FruitResponse)
	}

	return r0
}

// GetFruit provides a mock function with given fields: req
func (_m *Service) GetFruit(req *db.FruitRequest) db.Fruit {
	ret := _m.Called(req)

	var r0 db.Fruit
	if rf, ok := ret.Get(0).(func(*db.FruitRequest) db.Fruit); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(db.Fruit)
	}

	return r0
}

// GetFruits provides a mock function with given fields: req
func (_m *Service) GetFruits(req *db.FruitRequest) db.FruitResponse {
	ret := _m.Called(req)

	var r0 db.FruitResponse
	if rf, ok := ret.Get(0).(func(*db.FruitRequest) db.FruitResponse); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(db.FruitResponse)
	}

	return r0
}

// InsertFruit provides a mock function with given fields: req
func (_m *Service) InsertFruit(req *db.FruitRequest) db.FruitResponse {
	ret := _m.Called(req)

	var r0 db.FruitResponse
	if rf, ok := ret.Get(0).(func(*db.FruitRequest) db.FruitResponse); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(db.FruitResponse)
	}

	return r0
}