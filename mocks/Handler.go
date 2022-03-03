// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mysql "github.com/qq906907952/go-mysql/mysql"
	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// HandleFieldList provides a mock function with given fields: table, fieldWildcard
func (_m *Handler) HandleFieldList(table string, fieldWildcard string) ([]*mysql.Field, error) {
	ret := _m.Called(table, fieldWildcard)

	var r0 []*mysql.Field
	if rf, ok := ret.Get(0).(func(string, string) []*mysql.Field); ok {
		r0 = rf(table, fieldWildcard)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*mysql.Field)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(table, fieldWildcard)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleOtherCommand provides a mock function with given fields: cmd, data
func (_m *Handler) HandleOtherCommand(cmd byte, data []byte) error {
	ret := _m.Called(cmd, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(byte, []byte) error); ok {
		r0 = rf(cmd, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HandleQuery provides a mock function with given fields: query
func (_m *Handler) HandleQuery(query string) (*mysql.Result, error) {
	ret := _m.Called(query)

	var r0 *mysql.Result
	if rf, ok := ret.Get(0).(func(string) *mysql.Result); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mysql.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleStmtClose provides a mock function with given fields: context
func (_m *Handler) HandleStmtClose(context interface{}) error {
	ret := _m.Called(context)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(context)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HandleStmtExecute provides a mock function with given fields: context, query, args
func (_m *Handler) HandleStmtExecute(context interface{}, query string, args []interface{}) (*mysql.Result, error) {
	ret := _m.Called(context, query, args)

	var r0 *mysql.Result
	if rf, ok := ret.Get(0).(func(interface{}, string, []interface{}) *mysql.Result); ok {
		r0 = rf(context, query, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mysql.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, string, []interface{}) error); ok {
		r1 = rf(context, query, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleStmtPrepare provides a mock function with given fields: query
func (_m *Handler) HandleStmtPrepare(query string) (int, int, interface{}, error) {
	ret := _m.Called(query)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(query)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 interface{}
	if rf, ok := ret.Get(2).(func(string) interface{}); ok {
		r2 = rf(query)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(interface{})
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(string) error); ok {
		r3 = rf(query)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// UseDB provides a mock function with given fields: dbName
func (_m *Handler) UseDB(dbName string) error {
	ret := _m.Called(dbName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(dbName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
