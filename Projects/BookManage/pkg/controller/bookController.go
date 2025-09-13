package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rishi2600/Go-tutorial/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	fmt.Printf("Books: %v", res)
}
