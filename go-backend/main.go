package main

import (
	"context"
	"log"

	"github.com/mmvergara/biscord/go-backend/app"
)

func main() {

	app := app.NewApp(app.LoadConfig())

	err := app.Start(context.TODO())
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}