// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/kulinsky/guess_word/domain"
	mock "github.com/stretchr/testify/mock"
)

// WordLogic is an autogenerated mock type for the WordLogic type
type WordLogic struct {
	mock.Mock
}

// GetRandomWord provides a mock function with given fields: ctx
func (_m *WordLogic) GetRandomWord(ctx context.Context) (*domain.Word, error) {
	ret := _m.Called(ctx)

	var r0 *domain.Word
	if rf, ok := ret.Get(0).(func(context.Context) *domain.Word); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Word)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WordCreate provides a mock function with given fields: ctx, s
func (_m *WordLogic) WordCreate(ctx context.Context, s string) (*domain.Word, error) {
	ret := _m.Called(ctx, s)

	var r0 *domain.Word
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Word); ok {
		r0 = rf(ctx, s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Word)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}