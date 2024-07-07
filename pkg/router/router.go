package router

import (
	"net/http"

	"github.com/syshil/talks/internal/handlers"
)

// SetupRoutes initializes the server routes
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", handlers.ChatHandler)
	// Add more routes here
	return mux
}
