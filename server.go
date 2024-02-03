package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request", r)
	Notify(r)
}

func main() {
	if !setup() {
		return
	}
	http.HandleFunc("/", handler)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
