package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	fmt.Println("Starting webserver at: localhost:8000")
	// Nil means default instance of the webserver
	http.ListenAndServe(":8000", nil)
}