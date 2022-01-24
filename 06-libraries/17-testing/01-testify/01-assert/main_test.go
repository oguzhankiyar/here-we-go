package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet_ShouldReturnError_WhenRequestIsNil(t *testing.T) {
	// Arrange
	var request *Request

	// Act
	response := Get(request)

	// Assert
	assert.NotNil(t, response)
	assert.False(t, response.Status)
	assert.Empty(t, response.Data)
	assert.Equal(t, RequestCouldNotBeNil, response.Error)
}

func TestGet_ShouldReturnError_WhenRequestPageIsNotValid(t *testing.T) {
	pages := []int{0, -2}

	for _, page := range pages {
		t.Run(fmt.Sprintf("WhenPageIs_%v", page), func(t *testing.T) {
			// Arrange
			request := &Request{
				Page: page,
			}

			// Act
			response := Get(request)

			// Assert
			assert.NotNil(t, response)
			assert.False(t, response.Status)
			assert.Empty(t, response.Data)
			assert.Equal(t, RequestPageCouldNotBeLessThanOrEqualZero, response.Error)
		})
	}
}

func TestGet_ShouldReturnData_WhenRequestIsValid(t *testing.T) {
	pages := []int{1, 10}

	for _, page := range pages {
		t.Run(fmt.Sprintf("WhenPageIs%v", page), func(t *testing.T) {
			// Arrange
			request := &Request{
				Page: page,
			}

			// Act
			response := Get(request)

			// Assert
			assert.NotNil(t, response)
			assert.True(t, response.Status)
			assert.NotEmpty(t, response.Data)
			assert.Nil(t, response.Error)
		})
	}
}
