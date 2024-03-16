package models

import "math/big"

type FibonacciRequest struct {
	Number *big.Int `json:"number"`
}

type FibonacciResponse struct {
	Result *big.Int `json:"result"`
}
