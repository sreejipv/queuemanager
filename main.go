package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var queue = NewQueue()

func main() {
	http.HandleFunc("/enqueue", enqueueHandler)

	http.HandleFunc("/dequeue", dequeueHandler)
	http.HandleFunc("/list", listHandler)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func enqueueHandler(w http.ResponseWriter, r *http.Request) {
	var item struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	queue.Enqueue(item.Name)
	w.WriteHeader(http.StatusOK)
}

func dequeueHandler(w http.ResponseWriter, r *http.Request) {
	item, err := queue.Dequeue()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(item))
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	items := queue.List()
	json.NewEncoder(w).Encode(items)
}
