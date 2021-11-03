// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ArangoDB is an autogenerated mock type for the ArangoDB type
type ArangoDB struct {
	mock.Mock
}

// CreateDocument provides a mock function with given fields: ctx, col, doc
func (_m *ArangoDB) CreateDocument(ctx context.Context, col string, doc interface{}) error {
	ret := _m.Called(ctx, col, doc)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, col, doc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateDocuments provides a mock function with given fields: ctx, col, docs
func (_m *ArangoDB) CreateDocuments(ctx context.Context, col string, docs interface{}) error {
	ret := _m.Called(ctx, col, docs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, col, docs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DocumentExists provides a mock function with given fields: ctx, col, key
func (_m *ArangoDB) DocumentExists(ctx context.Context, col string, key string) (bool, error) {
	ret := _m.Called(ctx, col, key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, col, key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, col, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Query provides a mock function with given fields: ctx, query, binVars
func (_m *ArangoDB) Query(ctx context.Context, query string, binVars map[string]interface{}) (interface{}, error) {
	ret := _m.Called(ctx, query, binVars)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]interface{}) interface{}); ok {
		r0 = rf(ctx, query, binVars)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]interface{}) error); ok {
		r1 = rf(ctx, query, binVars)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadDocument provides a mock function with given fields: ctx, col, key, result
func (_m *ArangoDB) ReadDocument(ctx context.Context, col string, key string, result interface{}) error {
	ret := _m.Called(ctx, col, key, result)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, interface{}) error); ok {
		r0 = rf(ctx, col, key, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadDocuments provides a mock function with given fields: ctx, col, key, results
func (_m *ArangoDB) ReadDocuments(ctx context.Context, col string, key []string, results interface{}) error {
	ret := _m.Called(ctx, col, key, results)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, interface{}) error); ok {
		r0 = rf(ctx, col, key, results)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveDocument provides a mock function with given fields: ctx, col, key
func (_m *ArangoDB) RemoveDocument(ctx context.Context, col string, key string) error {
	ret := _m.Called(ctx, col, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, col, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDocument provides a mock function with given fields: ctx, col, key, doc
func (_m *ArangoDB) UpdateDocument(ctx context.Context, col string, key string, doc interface{}) error {
	ret := _m.Called(ctx, col, key, doc)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, interface{}) error); ok {
		r0 = rf(ctx, col, key, doc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
