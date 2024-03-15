package models

type FibonacciRequest struct {
	Number int64 `json:"number"`
}

type FibonacciResponse struct {
	Result int64 `json:"result"`
}
