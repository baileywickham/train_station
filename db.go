package main

import "database/sql"
import _ "github.com/lib/pq"

// single global db pointer for now
var db *sql.DB

func InitDB(connstr string) {
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		// db error unrecoverable at this point
		panic(err)
	}

	// db conn is defered til ping
	if err = db.Ping(); err != nil {
		panic(err)
	}
}
