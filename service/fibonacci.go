package service

import (
	"math"
	"math/big"

	"github.com/Hamedblue1381/bigfloat"
)

type FibonacciServicer interface {
	NextNumber(*big.Int) *big.Int
	PrevNumber(*big.Int) *big.Int
}

type FibonacciService struct {
}

// NewFibonacciService creates a new instance of FibonacciService.
func NewFibonacciService() FibonacciServicer {
	return &FibonacciService{}
}

// NextNumber calculates the next Fibonacci number based on the current number.
// It iterates through the Fibonacci sequence until it reaches or exceeds the current number.
func (fs *FibonacciService) NextNumber(current *big.Int) *big.Int {
	pos := fibonacciPosition(current)
	res := fibonacciNumber(pos.Add(pos, big.NewFloat(1)))
	return res
}

// PrevNumber calculates the previous Fibonacci number based on the current number.
// It iterates through the Fibonacci sequence in reverse until it finds the largest Fibonacci number less than the current number.
func (fs *FibonacciService) PrevNumber(current *big.Int) *big.Int {
	pos := fibonacciPosition(current)
	res := fibonacciNumber(pos.Sub(pos, big.NewFloat(1)))
	return res
}

func fibonacciNumber(position *big.Float) *big.Int {
	// Calculate phi
	phi := big.NewFloat(1).Add(big.NewFloat(1.0), new(big.Float).Sqrt(big.NewFloat(5.0)))
	phi.Quo(phi, big.NewFloat(2.0))

	// Apply Binet's formula: F(n) = (phi^n - (-phi)^-n) / sqrt(5)
	phiPowerN := bigfloat.Pow(phi, position)

	minusPhi := new(big.Float).Neg(phi)
	minusPos := new(big.Float).Neg(position)
	minusPhiPowerN := bigfloat.Pow(minusPhi, minusPos)
	fib := new(big.Float).Sub(phiPowerN, minusPhiPowerN)
	fib.Quo(fib, new(big.Float).Sqrt(big.NewFloat(5.0)))

	// Convert the result to a big.Int (rounded to nearest integer)
	fibInt := new(big.Int)
	fib.Int(fibInt)

	return fibInt
}

func fibonacciPosition(fib *big.Int) *big.Float {
	// Calculate (F(n) * sqrt(5) + 0.5)
	temp := new(big.Float).SetInt(fib)
	temp.Mul(temp, big.NewFloat(math.Sqrt(5)))
	temp.Add(temp, big.NewFloat(0.5))

	tempac, _ := temp.Float64()

	// Calculate log((F(n) * sqrt(5) + 0.5)) / log(phi)
	numerator := new(big.Float).SetFloat64(math.Log(tempac))

	denominator := new(big.Float).SetFloat64(math.Log(math.Phi))

	position := new(big.Float).Quo(numerator, denominator)

	return position
}

// Fibonacci calculates the Fibonacci number for a given input.
// It iterates through the Fibonacci sequence until it reaches or exceeds the input number.
func Fibonacci(x *big.Int) *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)

	for a.Cmp(x) <= 0 {
		a.Add(a, b)
		a, b = b, a
	}
	return a
}

// // Fibonacci calculates the Fibonacci number at position n.
// func Fibonacci(n float64) float64 {
// 	phi := (1 + math.Sqrt(5)) / 2
// 	return math.Round((math.Pow(phi, n) - math.Pow(-phi, -n)) / math.Sqrt(5))
// }

// // FibonacciPosition calculates the position of x in the Fibonacci sequence.
// func FibonacciPosition(x float64) float64 {
// 	phi := (1 + math.Sqrt(5)) / 2
// 	return math.Round(math.Log(x*math.Sqrt(5)+0.5) / math.Log(phi))
// }
