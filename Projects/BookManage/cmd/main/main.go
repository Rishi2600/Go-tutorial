package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Rishi2600/Go-tutorial/pkg/route"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()

	route.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	port := 8000
	fmt.Printf("server listening on port: %v \n", port)
	err := http.ListenAndServe("localhost:9010", r)

	if err != nil {
		log.Fatal(err)
	}
}
