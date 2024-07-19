package db

import (
	"log"

	"github.com/jackc/pgx"
	"github.com/mmvergara/biscord/go-backend/config"
)

func NewDb() *pgx.Conn {

	w := pgx.ConnConfig{
		Host:     config.Envs.PublicHost,
		Port:     config.Envs.DBPort,
		User:     config.Envs.DBUsername,
		Database: config.Envs.DBDatabase,
		Password: config.Envs.DBPassword,
	}
	conn, err := pgx.Connect(w)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return conn
}
