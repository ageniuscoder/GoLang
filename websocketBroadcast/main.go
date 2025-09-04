package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) //keep track of active clients

var mutex = sync.Mutex{}

var broadcast = make(chan []byte)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Error upgrading to websocket")
	}
	defer conn.Close()

	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	log.Println("New client added")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Client disconnected:", err)
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			break
		}
		log.Println("Message Recieved", msg)
		broadcast <- msg //non blocking
	}
}

func handleBroadcast() {
	for {
		msg := <-broadcast //blocking

		mutex.Lock()

		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			log.Println("Message broadcasted", string(msg))
			if err != nil {
				log.Println("Write error:", err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", handleWebsocket)
	go handleBroadcast()
	http.ListenAndServe(":8080", mux)
}
