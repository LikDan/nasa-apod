package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

var database *pgx.Conn

func init() {
	var err error
	database, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	createTables()
}

func createTables() {
	q, err := database.Query(context.Background(), `
CREATE TABLE IF NOT EXISTS apod_events
(
    id          SERIAL PRIMARY KEY,
    date        date UNIQUE,
    title       varchar,
    explanation text,
    picture     varchar
)
`)
	if err != nil {
		return
	}

	q.Close()
}
