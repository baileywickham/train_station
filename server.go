package main

import "errors"

import "log"

type Account struct {
	uuid    string // need to be generated by db?
	name    string
	balance int
}

var ErrUUIDConflict = errors.New("UUID already exists")
var ErrAcountBelowZ = errors.New("Account balance below zero")

var Accounts map[string]Account

func Create_account(uuid string, name string, balance int) (Account, error) {
	// check if account exists
	var account Account
	if _, ok := Accounts[uuid]; ok {
		// change error type?
		return account, ErrUUIDConflict
	}
	account = Account{uuid, name, balance}
	Accounts[uuid] = account
	return account, nil
}

func Add_balance(uuid string, amount int) error {
	account, err := account_by_uuid(uuid)
	if err != nil {
		return err
	}
	account.balance += amount
	Accounts[account.uuid] = account
	return nil
}

func account_by_uuid(uuid string) (Account, error) {
	var account Account
	if account, ok := Accounts[uuid]; ok {
		return account, nil
	}

	return account, errors.New("Account not found")
}
func print_account(uuid string) {
	account, err := account_by_uuid(uuid)
	if err != nil {
		return
	}
	println(account.uuid, account.name, account.balance)
}

func Charge_account(uuid string, amount int) error {
	account, err := account_by_uuid(uuid)
	if err != nil {
		return err
	}
	if account.balance-amount < 0 {
		return ErrAcountBelowZ
	}
	account.balance -= amount
	Accounts[account.uuid] = account

	return nil
}

func main() {
	runner()
	_, err := Create_account("1", "Bailey", 0)
	if err != nil {
		log.Fatal(err)
	}

	err = Add_balance("1", 10)
	if err != nil {
		log.Fatal(err)
	}

}
func init() {
	Accounts = make(map[string]Account)
}
