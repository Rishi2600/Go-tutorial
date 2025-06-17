package handlers

import (
	"github.com/Rishi2600/Go-tutorial/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {

	//Global Middleware to ignore slashes
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		//Middleware for /account route --authorization
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}

/*If a func name starts with a capital letter, it can be imported to other packages else it is a private func*/
