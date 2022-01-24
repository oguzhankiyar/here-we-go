package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (m *ServiceMock) GetCounts(length int) (map[string]int, error) {
	args := m.Called(length)

	var result map[string]int
	resultArg := args.Get(0)
	if resultArg != nil {
		result = resultArg.(map[string]int)
	}

	var err error
	errArg := args.Get(1)
	if errArg != nil {
		err = errArg.(error)
	}

	return result, err
}

func TestRun_ShouldReturnError_WhenIdIsNotSpecified(t *testing.T) {
	// Arrange
	serviceMock := &ServiceMock{}

	serviceMock.
		On("GetCounts", mock.Anything).
		Return(map[string]int{}, nil)

	// Act
	result, err := Run(serviceMock, "")

	// Assert
	assert.Equal(t, 0, result)
	assert.NotNil(t, err)

	serviceMock.AssertNotCalled(t, "GetCounts", mock.Anything)
}

func TestRun_ShouldReturnError_WhenServiceFails(t *testing.T) {
	// Arrange
	serviceMock := &ServiceMock{}

	serviceMock.
		On("GetCounts", mock.Anything).
		Return(nil, errors.New("some error occurred"))

	// Act
	result, err := Run(serviceMock, "MTk3NjIzNT")

	// Assert
	assert.Equal(t, 0, result)
	assert.NotNil(t, err)

	serviceMock.AssertCalled(t, "GetCounts", mock.Anything)
}

func TestRun_ShouldReturnError_WhenServiceReturnEmpty(t *testing.T) {
	// Arrange
	serviceMock := &ServiceMock{}

	serviceMock.
		On("GetCounts", mock.Anything).
		Return(make(map[string]int), nil)

	// Act
	result, err := Run(serviceMock, "MTk3NjIzNT")

	// Assert
	assert.Equal(t, 0, result)
	assert.NotNil(t, err)

	serviceMock.AssertCalled(t, "GetCounts", mock.Anything)
}

func TestRun_ShouldReturnError_WhenServiceDoesNotContainSpecifiedId(t *testing.T) {
	// Arrange
	serviceMock := &ServiceMock{}

	serviceMock.
		On("GetCounts", mock.Anything).
		Return(map[string]int{
			"MTk3Nj3f6B": 100,
		}, nil)

	// Act
	result, err := Run(serviceMock, "MTk3NjIzNT")

	// Assert
	assert.Equal(t, 0, result)
	assert.NotNil(t, err)

	serviceMock.AssertCalled(t, "GetCounts", mock.Anything)
}

func TestRun_ShouldReturnCount_WhenServiceContainsSpecifiedId(t *testing.T) {
	// Arrange
	serviceMock := &ServiceMock{}

	var length int
	serviceMock.
		On("GetCounts", mock.Anything).
		Run(func(args mock.Arguments) {
			length = args.Int(0)
		}).
		Return(map[string]int{
			"MTk3NjIzNT": 100,
		}, nil)

	// Act
	result, err := Run(serviceMock, "MTk3NjIzNT")

	// Assert
	assert.Equal(t, 100, result)
	assert.Nil(t, err)
	assert.Equal(t, 10, length)

	serviceMock.AssertCalled(t, "GetCounts", mock.Anything)
}
