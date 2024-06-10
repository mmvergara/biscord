package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes() {

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	// Defining routes
	router.Get("/",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello World!"))
	})

	router.Route("/v1", a.loadV1Routes)

	// End of defining routes

	// Assign the router to the app
	a.router = router

}




// Load v1 routes
func (a *App) loadV1Routes(router chi.Router) {

	
	
	// router.Post("/register")
	// router.Post("/login")
}