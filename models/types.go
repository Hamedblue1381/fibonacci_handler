package models

type FibonacciRequest struct {
	Number int `json:"number"`
}

type FibonacciResponse struct {
	Result int `json:"result"`
}
