package route

import (
	"github.com/Rishi2600/Go-tutorial/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controller.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookid}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookid}", controller.updateBook).Methods("PUT")
	router.HandleFunc("/book/{bookid}", controller.DeleteBook).Methods("DELETE")
}

//the purpose behind adding all the routes in a func other than main and assigning it to a variable,
//is just to add all the values of the routes to that particular variable and thus the var can be used further in different
//packages or function (including main).
