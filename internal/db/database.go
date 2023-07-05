package database

import (
	"database/sql"
	"log"
)

type DB struct {
	Connection *sql.DB
}

func NewPool(connectionString string) *DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln("err")
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		log.Fatalln("err")
	}

	return &DB{
		Connection: db,
	}
}
