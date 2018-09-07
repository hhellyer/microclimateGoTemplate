package main

import (
	"io"
	"net/http"
)

// Respond to call on endpoint
func respond(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from microclimate")
}

func main() {
	http.HandleFunc("/", respond)
	http.ListenAndServe(":8000", nil)
}
