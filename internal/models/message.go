package models

// Message represents a chat message from a user.
type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}
