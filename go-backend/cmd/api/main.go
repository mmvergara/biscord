package main

import (
	"context"
	"log"
)

func main() {
	cfg := LoadConfig()
	
	app := NewApp(cfg)

	err := app.Start(context.TODO())
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}