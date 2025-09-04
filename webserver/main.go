package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	Name string `json:"name"`
}

var mu sync.RWMutex

var cacheUser = make(map[int]User)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello mangal server")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	mu.Lock()
	cacheUser[len(cacheUser)+1] = user
	mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	mu.RLock()
	user, ok := cacheUser[idInt]
	mu.RUnlock()

	if !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	w.Header().Set("content-type", "application/json")
	data, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := cacheUser[id]; !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	mu.Lock()
	delete(cacheUser, id)
	mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlerFunc)

	mux.HandleFunc("POST /users", createUser)

	mux.HandleFunc("GET /users/{id}", getUser)

	mux.HandleFunc("DELETE /users/{id}", deleteUser)

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: mux,
	}
	slog.Info("server started")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
