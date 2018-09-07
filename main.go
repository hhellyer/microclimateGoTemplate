package main

import (
	"io"
	"net/http"
)

// Respond to call on endpoint
func respond(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from your Go sample application running in Microclimate!")
}

func main() {
	http.HandleFunc("/", respond)
	http.ListenAndServe(":8000", nil)
}
