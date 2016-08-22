package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", get)
	http.ListenAndServe(":8080", nil)
}

func get(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello http!")
}
