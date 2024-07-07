package models

import "github.com/gorilla/websocket"

// Session represents a chat session between two users.
type Session struct {
	User1 *websocket.Conn
	User2 *websocket.Conn
}
