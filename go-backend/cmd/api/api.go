package api

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jackc/pgx"
	"github.com/labstack/echo/v4"
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
	e := echo.New()

	e.Use(config.CorsOptions)
	// e.Use(middleware.Logger())

	userRepo := repo.NewUserRepo(s.db)

	authHandler := auth.NewHandler(userRepo)

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	v1 := e.Group("/v1")
	v1.GET("/", func(c echo.Context) error {
		playground.Handler("GraphQL playground", "/v1/graph").ServeHTTP(c.Response(), c.Request())
		return nil
	})
	v1.POST("/auth/sign-in", authHandler.SignIn)
	v1.POST("/auth/sign-up", authHandler.SignUp)
	// v1.POST("/auth/sign-out", s.Refresh)

	srv := handler.NewDefaultServer(graph_generated.NewExecutableSchema(graph_generated.Config{Resolvers: &graph_resolvers.Resolver{}}))

	v1.Use() // auth middleware
	v1.Use() // data loader middleware
	v1.POST("/graph", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	}) // graphql endpoint

	e.Logger.Fatal(e.Start(s.addr))

	e.Start(s.addr)
}
