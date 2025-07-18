package main

import (
	"fmt"
	"net/http"
)

var port int = 3000

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello from port: ", port)
}

func main() {
	http.HandleFunc("/hello", hello)

	fmt.Printf("Server listening on port: %v...\n", port)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
