package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error Upgrading", err)
	}
	defer conn.Close()
	log.Println("Client connected")

	for {
		msgtype, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Recieved msg %s", msg)

		time.Sleep(3 * time.Second)

		err = conn.WriteMessage(msgtype, msg)

		if err != nil {
			log.Println("write error", err)
			break
		}

		log.Printf("Send msg %s", msg)
	}
}

func main() {

	http.HandleFunc("/ws", websocketHandler)
	fmt.Println("Connected to the server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error connecting to server")
	}

}
