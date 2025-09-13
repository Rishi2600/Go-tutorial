package route

import (
	"github.com/Rishi2600/Go-tutorial/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controller.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookid}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookid}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookid}", controller.DeleteBook).Methods("DELETE")
}
