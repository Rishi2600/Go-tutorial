package route

import (
	"github.com/Rishi2600/Go-tutorial/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
}
