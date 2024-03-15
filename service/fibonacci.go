package service

import "math"

type FibonacciServicer interface {
	NextNumber(int64) int64
	PrevNumber(int64) int64
}

type FibonacciService struct {
}

// NewFibonacciService creates a new instance of FibonacciService.
func NewFibonacciService() FibonacciServicer {
	return &FibonacciService{}
}

// NextNumber calculates the next Fibonacci number based on the current number.
// If the current number is 0, it returns 1.
func (fs *FibonacciService) NextNumber(current int64) int64 {
	if current == 0 {
		return 1
	}
	pos := FibonacciPosition(float64(current))
	next := Fibonacci(pos + 1)
	return int64(next)
}

// PrevNumber calculates the previous Fibonacci number based on the current number.
// If the current number is less than or equal to 1, it returns current - 1.
func (fs *FibonacciService) PrevNumber(current int64) int64 {
	if current <= 1 {
		return current - 1
	}
	pos := FibonacciPosition(float64(current))
	prev := Fibonacci(pos - 1)
	return int64(prev)
}

// Fibonacci calculates the Fibonacci number at position n.
func Fibonacci(n float64) float64 {
	phi := (1 + math.Sqrt(5)) / 2
	return math.Round((math.Pow(phi, n) - math.Pow(-phi, -n)) / math.Sqrt(5))
}

// FibonacciPosition calculates the position of x in the Fibonacci sequence.
func FibonacciPosition(x float64) float64 {
	phi := (1 + math.Sqrt(5)) / 2
	return math.Round(math.Log(x*math.Sqrt(5)+0.5) / math.Log(phi))
}
