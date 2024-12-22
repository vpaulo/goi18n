package main

import (
	"fmt"
	"goi18n/internal/server"
	"log"
)

func main() {
	server := server.NewServer()

	log.Printf("Listening on %s...", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
