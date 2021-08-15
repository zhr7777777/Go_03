package main

import (
	"fmt"
	"net/http"
)

func hello(rw http.ResponseWriter, r *http.Request) {
	fmt.Println()
	fmt.Fprintf(rw, "Hello\n")
}

func main() {
	http.HandleFunc("/api/hello", func(rw http.ResponseWriter, r *http.Request) {
		LimitRequest(rw, r, hello, 2)
	})

	http.ListenAndServe(":8080", nil)
}
