package main

import "errors"
import "log"

type Account struct {
	uuid    int // need to be generated by db?
	name    string
	balance int
}

var ErrUUIDConflict = errors.New("UUID already exists")
var ErrAcountBelowZ = errors.New("Account balance below zero")
var ErrAccountNotFound = errors.New("Account not found")

//var Accounts map[int]Account

func Create_account(name string, balance int) (Account, error) {
	// check if account exists
	id, err := Create_account_db(name, balance)
	if err != nil {
		panic(err)
	}
	account := Account{id, name, balance}

	return account, nil
}

func Add_balance(uuid int, amount int) error {
	account, err := account_by_uuid(uuid)
	if err != nil {
		return err
	}
	account.balance += amount
	_ = update_account_db(account)
	return nil
}

func account_by_uuid(uuid int) (Account, error) {
	account, err := account_by_uuid_db(uuid)
	if err != nil {
		return account, err
	}

	return account, nil
}

func print_account(uuid int) {
	account, err := account_by_uuid(uuid)
	if err != nil {
		log.Println(err)
		return
	}
	println(account.uuid, account.name, account.balance)
}

func Charge_account(uuid int, amount int) error {
	account, err := account_by_uuid(uuid)
	if err != nil {
		return err
	}
	if account.balance-amount < 0 {
		return ErrAcountBelowZ
	}
	account.balance -= amount
	// If account not found, will fail earlier
	_ = update_account_db(account)

	return nil
}
func Delete_user(uuid int) error {
	err := delete_user_db(uuid)
	return err
}

func main() {
	runner()
}
func init() {
	InitDB()
}
