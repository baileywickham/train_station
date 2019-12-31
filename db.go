package main

import "database/sql"
import _ "github.com/lib/pq"
import "time"

// single global db pointer for now
var db *sql.DB
var connStr = "user=postgres password=postgres dbname=station sslmode=disable"
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
	err := db.QueryRow("SELECT * FROM accounts WHERE id = $1", uuid).Scan(&account.UUID, &account.Name, &account.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return account, ErrAccountNotFound
		}
		panic(err)
	}
	return account, nil
}

func Create_account_db(name string, balance int) int {
	var id int
	err := db.QueryRow(
		`INSERT INTO accounts (name, balance)
		VALUES ($1, $2)
		RETURNING id`, name, balance).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}

func update_account_db(account Account) {
	_, err := db.Exec(`UPDATE accounts
				SET name=$1, balance=$2
				WHERE id=$3`, account.Name, account.Balance, account.UUID)
	if err != nil {
		// Handle errors in func
		panic(err)
	}
}

func delete_user_db(uuid int) error {
	_, err := db.Exec(`DELETE FROM accounts WHERE id=$1`, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrAccountNotFound
		}
		panic(err)
	}
	return nil
}

func get_all_accounts_db() []Account {
	// default size of 10
	accounts := make([]Account, 0)
	rows, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var a Account
		err = rows.Scan(&a.UUID, &a.Name, &a.Balance)
		if err != nil {
			panic(err)
		}
		accounts = append(accounts, a)
	}
	return accounts
}

func updated_accounts(ac <-chan Account) {
	for {
		select {
		case account, ok := <-ac:
			if ok == false {
				return
			}
			account.channged = false
			go update_account_db(account)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
