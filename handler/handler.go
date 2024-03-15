package handler

import (
	"net/http"

	"github.com/Hamedblue1381/fibonacci_handler/models"
	"github.com/Hamedblue1381/fibonacci_handler/service"
)

func NextNumberHandler(w http.ResponseWriter, r *http.Request) {
	var req models.FibonacciRequest
	if err := service.ParseRequest(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fs := service.NewFibonacciService()
	result := fs.NextNumber(req.Number)

	response := models.FibonacciResponse{Result: result}

	if err := service.EncodeResponse(w, r, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PrevNumberHandler(w http.ResponseWriter, r *http.Request) {
	var req models.FibonacciRequest
	if err := service.ParseRequest(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fs := service.NewFibonacciService()
	result := fs.PrevNumber(req.Number)

	response := models.FibonacciResponse{Result: result}

	if err := service.EncodeResponse(w, r, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
