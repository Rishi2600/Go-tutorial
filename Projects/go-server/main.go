package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 route not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "request method is wrong", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Fprintf(w, "parse form err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request seccuessfull")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name: %v \n", name)
	fmt.Fprintf(w, "Address: %v \n", address)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	port := 8080
	fmt.Printf("server listening on port: %v \n", port)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
