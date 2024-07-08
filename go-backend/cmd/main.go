package main

import (
	"context"

	"github.com/mmvergara/biscord/go-backend/cmd/api"
	db "github.com/mmvergara/biscord/go-backend/internal/database"
)

func main() {
	ctx := context.Background()
	conn := db.NewDb()
	err := conn.Ping(ctx)
	if err != nil {
		panic(err)
	}
	// get all tables
	rows, err := conn.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	println("Tables:")
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			panic(err)
		}
		println(tableName)
	}

	api.NewServer("localhost:8080", conn).Run()

}
