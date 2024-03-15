package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciService_NextNumber(t *testing.T) {
	fs := FibonacciService{}

	assert.Equal(t, 1, fs.NextNumber(0))
	assert.Equal(t, 2, fs.NextNumber(1))
	assert.Equal(t, 5, fs.NextNumber(3))
	assert.Equal(t, 8, fs.NextNumber(5))
}

func TestFibonacciService_PrevNumber(t *testing.T) {
	fs := FibonacciService{}

	assert.Equal(t, 0, fs.PrevNumber(1))
	assert.Equal(t, 1, fs.PrevNumber(2))
	assert.Equal(t, 3, fs.PrevNumber(5))
	assert.Equal(t, 5, fs.PrevNumber(8))
}
