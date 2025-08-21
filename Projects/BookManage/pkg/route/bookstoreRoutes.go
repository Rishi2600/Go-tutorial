package route

import (
	"github.com/Rishi2600/Go-tutorial/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func() {
	var router *mux.Router = mux.NewRouter()

	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
}
