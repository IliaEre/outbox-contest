package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"

	"order-service/config"
)

func NewConnection(cfg config.Config) *sql.DB {
	dbc := cfg.Database
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbc.Username, dbc.Database, dbc.Password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
