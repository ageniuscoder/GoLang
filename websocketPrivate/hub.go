package main

import (
	"encoding/json" // ✅ UPDATED
	"log"
)

// ✅ UPDATED: Added Sender field and changed Content to string
type Message struct {
	Sender    string `json:"sender"` // ✅ UPDATED
	Recipient string `json:"recipient"`
	Content   string `json:"content"` // ✅ UPDATED (was []byte)
}

type Hub struct {
	clients        map[string]*Client
	privateMessage chan Message
	register       chan *Client
	unregister     chan *Client
}

func newHub() *Hub {
	return &Hub{
		privateMessage: make(chan Message),
		register:       make(chan *Client),
		unregister:     make(chan *Client),
		clients:        make(map[string]*Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			log.Printf("Client registered: %s", client.userID)
			h.clients[client.userID] = client

		case client := <-h.unregister:
			log.Printf("Client unregistered: %s", client.userID)
			if _, ok := h.clients[client.userID]; ok {
				delete(h.clients, client.userID)
				close(client.send)
			}

		case message := <-h.privateMessage:
			if recipient, ok := h.clients[message.Recipient]; ok {
				// ✅ UPDATED: Marshal sender + content into JSON
				jsonResponse, err := json.Marshal(map[string]string{
					"sender":  message.Sender,
					"content": message.Content,
				})
				if err != nil {
					log.Printf("Error marshalling response: %v", err)
					continue
				}

				select {
				case recipient.send <- jsonResponse: // ✅ UPDATED
					log.Printf("Message sent from %s to %s", message.Sender, message.Recipient)
				default:
					close(recipient.send)
					delete(h.clients, message.Recipient)
					log.Printf("Failed to send message to %s", message.Recipient)
				}
			} else {
				log.Printf("Recipient %s not found", message.Recipient)
			}
		}
	}
}
