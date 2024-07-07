package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/syshil/talks/pkg/router"
)

func main() {
	fmt.Println("Starting the game server on :8080")
	mux := router.SetupRoutes()
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
