package service

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciService_NextNumber(t *testing.T) {
	fs := FibonacciService{}
	numStr299 := "137347080577163115432025771710279131845700275212767467264610201"

	num299 := new(big.Int)
	num299.SetString(numStr299, 10)

	numStr300 := "222232244629422676106398124069050176556246085874450435841982464"

	num300 := new(big.Int)
	num300.SetString(numStr300, 10)

	assert.Equal(t, big.NewInt(55), fs.NextNumber(big.NewInt(34)))
	assert.Equal(t, num300, fs.NextNumber(num299))
	assert.Equal(t, big.NewInt(8), fs.NextNumber(big.NewInt(5)))
}

func TestFibonacciService_PrevNumber(t *testing.T) {
	fs := FibonacciService{}

	numStr299 := "137347080577163115432025771710279131845700275212767467264610201"

	num299 := new(big.Int)
	num299.SetString(numStr299, 10)

	numStr298 := "84885164052258178774158161262034422874929272117503685216436224"

	num298 := new(big.Int)
	num298.SetString(numStr298, 10)

	assert.Equal(t, big.NewInt(34), fs.PrevNumber(big.NewInt(55)))
	assert.Equal(t, num298, fs.PrevNumber(num299))
	assert.Equal(t, big.NewInt(3), fs.PrevNumber(big.NewInt(5)))
	assert.Equal(t, big.NewInt(5), fs.PrevNumber(big.NewInt(8)))
}
