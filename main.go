package main

import (
	"log"

	"github.com/Hamedblue1381/fibonacci_handler/server"
)

func main() {
	srv := server.SetupServer()

	log.Println("Server starting on port 8080...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
