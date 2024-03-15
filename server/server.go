package server

import (
	"net/http"

	"github.com/Hamedblue1381/fibonacci_handler/handler"
)

func SetupServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/next", handler.NextNumberHandler)
	mux.HandleFunc("/prev", handler.PrevNumberHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return server
}
