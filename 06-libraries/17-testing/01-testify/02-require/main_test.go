package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSomething(t *testing.T) {
	a := ""
	b := "Gopher"

	// The test will continue even if this fails
	println("Checking the equality between a and b")
	assert.Equal(t, a, b)

	// The test will be stopped after this fails
	println("Checking the a is not empty")
	require.NotEmpty(t, a)

	// This below lines will not be executed
	println("Checking the b is not empty")
	require.NotEmpty(t, b)
}
