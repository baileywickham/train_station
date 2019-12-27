package main

import "errors"

type Account struct {
	uuid    int // need to be generated by db?
	name    string
	balance int
}

var ErrUUIDConflict = errors.New("UUID already exists")
var ErrAcountBelowZ = errors.New("Account balance below zero")
var ErrAccountNotFound = errors.New("Account not found")

//var Accounts map[int]Account

func Create_account(name string, balance int) Account {
	// check if account exists
	id := Create_account_db(name, balance)
	account := Account{id, name, balance}

	return account
}

func (account *Account) Add_balance(amount int) {
	account.balance += amount
	account.update()
}

func (account *Account) update() {
	update_account_db(*account)
}

func account_by_uuid(uuid int) (Account, error) {
	account, err := account_by_uuid_db(uuid)
	if err != nil {
		return account, err
	}

	return account, nil
}

func (account *Account) print_account() {
	println(account.uuid, account.name, account.balance)
}

func (account *Account) Charge_account(amount int) error {
	if account.balance-amount < 0 {
		return ErrAcountBelowZ
	}
	account.balance -= amount
	account.update()
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
