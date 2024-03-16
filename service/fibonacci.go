package service

import "math/big"

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
	a := big.NewInt(0)
	b := big.NewInt(1)

	for a.Cmp(current) <= 0 {
		a.Add(a, b)
		a, b = b, a
	}

	return a
}

// PrevNumber calculates the previous Fibonacci number based on the current number.
// It iterates through the Fibonacci sequence in reverse until it finds the largest Fibonacci number less than the current number.
func (fs *FibonacciService) PrevNumber(current *big.Int) *big.Int {
	if current.Cmp(big.NewInt(0)) <= 0 {
		return big.NewInt(0)
	}
	a := big.NewInt(0)
	b := big.NewInt(1)
	c := new(big.Int).Set(current)

	for Fibonacci(a).Cmp(c) < 0 {
		a.Add(a, b)
		a, b = b, a
	}
	return a
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
