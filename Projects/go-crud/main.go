package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	fmt.Println(movies)
	movies = append(movies, Movie{ID: "1", Isbn: "01", Title: "first movie", Director: &Director{FirstName: "X", LastName: "Nolan"}})
	movies = append(movies, Movie{ID: "2", Isbn: "02", Title: "second movie", Director: &Director{FirstName: "Y", LastName: "Kashyap"}})
	fmt.Println(movies)

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	port := 8000
	fmt.Printf("server listening on port: %v \n", port)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
