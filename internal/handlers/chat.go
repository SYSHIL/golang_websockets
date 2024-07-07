package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/syshil/talks/internal/services"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ChatHandler called")
	userName := r.URL.Query().Get("user")
	if userName == "" {
		http.Error(w, "Missing user query parameter", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not establish WebSocket connection", http.StatusInternalServerError)
		return
	}

	services.JoinChat(userName, conn)
}
