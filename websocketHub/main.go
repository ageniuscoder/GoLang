package main

import (
	"log"
	"log/slog"
	"net/http"
)

var hub = newHub() //new hub instance

func handleWs(w http.ResponseWriter, r *http.Request) {
	serveWs(hub, w, r)
}

func main() {
	go hub.run()

	http.HandleFunc("/ws", handleWs)
	slog.Info("server started at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
