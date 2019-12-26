package main

import "database/sql"
import _ "github.com/lib/pq"

// single global db pointer for now
var db *sql.DB
var connStr = "user=y dbname=station sslmode=verify-full"

func InitDB() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		// db error unrecoverable at this point
		panic(err)
	}

	// db conn is defered til ping
	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func account_by_uuid_db(uuid string) (Account, error) {
	var account Account
	rows, err := db.Query("SELECT * FROM station WHERE uuid = ?", uuid)
	if err != nil {
		return account, err
	}
	defer rows.Close()
	err = rows.Scan(&account.uuid, &account.name, &account.balance)
	if err != nil {
		panic(err)
	}
	return account, nil
}

func Create_account_db(name string, balance int) (int, error) {
	res, err := db.Exec(
		`INSERT INTO accounts (name, balance)
		VALUES ($1, $2)
		RETURNING id`, name, balance)
	if err != nil {
		panic(err)
	}
	ret, err := res.LastInsertId()
	return int(ret), err
}
