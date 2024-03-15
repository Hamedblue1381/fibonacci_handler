package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciService_NextNumber(t *testing.T) {
	fs := FibonacciService{}
	fmt.Println(fs.NextNumber(int64(0)))

	assert.Equal(t, int64(1), fs.NextNumber(int64(0)))
	assert.Equal(t, int64(2), fs.NextNumber(int64(1)))
	assert.Equal(t, int64(5), fs.NextNumber(int64(3)))
	assert.Equal(t, int64(8), fs.NextNumber(int64(5)))
}

func TestFibonacciService_PrevNumber(t *testing.T) {
	fs := FibonacciService{}

	assert.Equal(t, int64(0), fs.PrevNumber(int64(1)))
	assert.Equal(t, int64(1), fs.PrevNumber(int64(2)))
	assert.Equal(t, int64(3), fs.PrevNumber(int64(5)))
	assert.Equal(t, int64(5), fs.PrevNumber(int64(8)))
}
