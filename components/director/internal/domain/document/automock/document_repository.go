// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/director/internal/model"

// DocumentRepository is an autogenerated mock type for the DocumentRepository type
type DocumentRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: item
func (_m *DocumentRepository) Create(item *model.Document) error {
	ret := _m.Called(item)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Document) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: item
func (_m *DocumentRepository) Delete(item *model.Document) error {
	ret := _m.Called(item)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Document) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *DocumentRepository) GetByID(id string) (*model.Document, error) {
	ret := _m.Called(id)

	var r0 *model.Document
	if rf, ok := ret.Get(0).(func(string) *model.Document); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Document)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListByApplicationID provides a mock function with given fields: applicationID, pageSize, cursor
func (_m *DocumentRepository) ListByApplicationID(applicationID string, pageSize *int, cursor *string) (*model.DocumentPage, error) {
	ret := _m.Called(applicationID, pageSize, cursor)

	var r0 *model.DocumentPage
	if rf, ok := ret.Get(0).(func(string, *int, *string) *model.DocumentPage); ok {
		r0 = rf(applicationID, pageSize, cursor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DocumentPage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *int, *string) error); ok {
		r1 = rf(applicationID, pageSize, cursor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
