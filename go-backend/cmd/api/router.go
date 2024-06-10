package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	// Defining routes

	// Load v1 routes
	router.Route("/v1", a.loadV1Routes)

	// End of defining routes

	a.router = router

}


func (a *App) loadV1Routes(router chi.Router) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("V1 API Ready!")) 
	})

	router.Route("/auth", a.loadV1AuthRoutes)


}

func (a *App) loadV1AuthRoutes(router chi.Router) {

}