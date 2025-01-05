package api

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx"
	"github.com/mmvergara/biscord/go-backend/config"
	"github.com/mmvergara/biscord/go-backend/services/auth"
	graph_generated "github.com/mmvergara/biscord/go-backend/services/graphql/generated"
	graph_resolvers "github.com/mmvergara/biscord/go-backend/services/graphql/resolvers"
	"github.com/mmvergara/biscord/go-backend/services/repo"
)

type Server struct {
	addr string
	db   *pgx.Conn
}

func NewServer(addr string, db *pgx.Conn) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Run() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(config.CorsOptions)

	userRepo := repo.NewUserRepo(s.db)
	authHandler := auth.NewHandler(userRepo)
	srv := handler.NewDefaultServer(graph_generated.NewExecutableSchema(graph_generated.Config{Resolvers: &graph_resolvers.Resolver{}}))

	r.Get("/", s.HelloWorld)
	r.Route("/v1", func(r chi.Router) {
		r.Get("/", playground.Handler("GraphQL playground", "/v1/graph").ServeHTTP)

		r.Post("/auth/sign-in", authHandler.SignInHandler)
		r.Post("/auth/sign-up", authHandler.SignUpHandler)
		r.Post("/auth/sign-out", authHandler.SignOutHandler)

		r.Group(func(r chi.Router) {
			r.Use(authHandler.AuthMiddleware())
			// r.Use(DataLoaderMiddleware)

			r.Post("/auth/test", authHandler.AuthTestHandler)
			r.Post("/graph", func(w http.ResponseWriter, r *http.Request) {
				srv.ServeHTTP(w, r)
			})

		})
	})

	log.Printf("Server starting on %s", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, r))
}

func (s *Server) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
