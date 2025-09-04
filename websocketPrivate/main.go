package main

import (
	"log"
	"log/slog"
	"net/http"
)

func main() {
	hub := newHub()
	go hub.run()

	http.HandleFunc("/ws/{userid}", func(w http.ResponseWriter, r *http.Request) {
		//ex ws/mangal
		userID := r.PathValue("userid")
		if userID == "" {
			http.Error(w, "User ID not provided in URL path", http.StatusBadRequest)
			return
		}
		serveWs(hub, w, r, userID)
	})
	slog.Info("server started at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
