package main

import (
	"fmt"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()
	fmt.Println(r)

	log.Info(r)

}
