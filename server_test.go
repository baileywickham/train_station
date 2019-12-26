package main

import "errors"
import "testing"

var account Account

func TestCreate_account(t *testing.T) {
	var err error
	account, err = Create_account("test", 0)
	// only verify uuid cause im lazy
	if err != nil || account.name != "test" || account.balance != 0 {
		t.Log(err)
		t.Error("Account not created")
	}
}

func TestAdd_balance(t *testing.T) {
	err := Add_balance(account.uuid, 1)
	if err != nil {
		t.Log(err)
		t.Error("Balance not increased")
	}
	a, err := account_by_uuid(account.uuid)
	if err != nil || a.balance != 1 {
		t.Error("Balance not increased")
	}
}

func TestCharge_account(t *testing.T) {
	err := Charge_account(account.uuid, 1)
	a := account_by_uuid
	if err != nil || Accounts["1"].balance != 0 {
		t.Log(err)
		t.Error("Balance not decreased")
	}
	err = Charge_account(account.uuid, 10)
	if !errors.Is(ErrAcountBelowZ, err) {
		t.Error("Balance decreased below zero")
	}

}
