package service

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciService_NextNumber(t *testing.T) {
	fs := FibonacciService{}

	assert.Equal(t, big.NewInt(55), fs.NextNumber(big.NewInt(34)))
	assert.Equal(t, big.NewInt(3), fs.NextNumber(big.NewInt(2)))
	assert.Equal(t, big.NewInt(8), fs.NextNumber(big.NewInt(5)))
}

func TestFibonacciService_PrevNumber(t *testing.T) {
	fs := FibonacciService{}

	assert.Equal(t, big.NewInt(34), fs.PrevNumber(big.NewInt(55)))
	assert.Equal(t, big.NewInt(1), fs.PrevNumber(big.NewInt(2)))
	assert.Equal(t, big.NewInt(3), fs.PrevNumber(big.NewInt(5)))
	assert.Equal(t, big.NewInt(5), fs.PrevNumber(big.NewInt(8)))
}
