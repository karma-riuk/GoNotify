package main

import (
	"fmt"
	"io"
	"net/http"
)

func foo() string {
	return "Hello, World!"
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("foo says: %s", foo()))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
