package service

type FibonacciServicer interface {
	NextNumber(int) int
	PrevNumber(int) int
}

type FibonacciService struct {
}

func NewFibonacciService() FibonacciServicer {
	return &FibonacciService{}
}
func (fs *FibonacciService) NextNumber(current int) int {
	n := 0
	for {
		if fibonacci(n) > current {
			return fibonacci(n)
		}
		n++
	}
}

func (fs *FibonacciService) PrevNumber(current int) int {
	n := 0
	for {
		if fibonacci(n) >= current {
			return fibonacci(n - 1)
		}
		n++
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
