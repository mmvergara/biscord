package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	router http.Handler
	mdb *mongo.Client
	config Config
}

func NewApp(config Config) *App {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoDBAddr))
	mdb := mongoClient.Database("biscord")


	if err != nil {
		panic(err)
	}


	app := &App{
		mdb: mongoClient,
		config: config,
	}

	app.loadRoutes()

	
	return app

}

func (a *App) Start (ctx context.Context) error {
	server := &http.Server{
		Addr: "localhost:8080",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("server error: %w", err)
	}
	log.Println("Server started on localhost:8080")
	return nil
}