package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func openDB(dsn string) (*sql.DB, error) {
	// open db connection
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// check db is alive
	err = db.Ping()

	if err != nil {
		return nil, err
	}
	return db, nil
}

func (app *application) connectToDB() (*sql.DB, error) {
	connection, err := openDB(app.DSN)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Postgres!")

	return connection, nil
}
