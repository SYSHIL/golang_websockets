package services

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"

	"github.com/syshil/talks/internal/models"
)

var (
	sessions     = make(map[string]*models.Session)
	waitingUsers = make(map[string]*websocket.Conn)
	mu           sync.Mutex
)

func JoinChat(userName string, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()

	if len(waitingUsers) > 0 {
		for name, otherConn := range waitingUsers {
			delete(waitingUsers, name)
			session := &models.Session{User1: otherConn, User2: conn}
			sessions[name] = session
			sessions[userName] = session
			go handleMessages(session.User1, session.User2)
			go handleMessages(session.User2, session.User1)
			return
		}
	}

	waitingUsers[userName] = conn
}

func handleMessages(conn1, conn2 *websocket.Conn) {
	defer func() {
		conn1.Close()
		conn2.Close()
	}()

	for {
		_, msg, err := conn1.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		message := models.Message{Sender: "User", Content: string(msg)}
		err = conn2.WriteJSON(message)
		if err != nil {
			log.Println("Error sending message:", err)
			return
		}
	}
}
