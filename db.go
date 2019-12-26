package main

import "database/sql"
import _ "github.com/lib/pq"

// single global db pointer for now
var db *sql.DB
var connStr = "user=postgres password=testing dbname=station"
var table_name = "accounts"

func InitDB() {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		// db error unrecoverable at this point
		panic(err)
	}

	// db conn is defered til ping
	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func account_by_uuid_db(uuid int) (Account, error) {
	var account Account
	err := db.QueryRow("SELECT * FROM accounts WHERE id = $1", uuid).Scan(&account.uuid, &account.name, &account.balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return account, ErrAccountNotFound
		}
		panic(err)
	}
	return account, nil
}

func Create_account_db(name string, balance int) (int, error) {
	var id int
	err := db.QueryRow(
		`INSERT INTO accounts (name, balance)
		VALUES ($1, $2)
		RETURNING id`, name, balance).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id, nil
}

func update_account_db(account Account) error {
	_, err := db.Exec(`UPDATE $1
				SET name=$2, balance=$3
				WHERE id=$4`, table_name, account.name, account.balance, account.uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrAccountNotFound
		}
		panic(err)
	}
	return nil
}

func delete_user_db(uuid int) error {
	_, err := db.Exec(`DELETE FROM $1 WHERE id=$2`, table_name, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrAccountNotFound
		}
		panic(err)
	}
	return nil
}
