package main

import "errors"
import "testing"

func TestCreate_account(t *testing.T) {
	account, err := Create_account("1", "test", 0)
	// only verify uuid cause im lazy
	if err != nil || account.uuid != "1" {
		t.Log(err)
		t.Error("Account not created")
	}
}

func TestAdd_balance(t *testing.T) {
	err := Add_balance("1", 1)
	if err != nil || Accounts["1"].balance != 1 {
		t.Log(err)
		t.Error("Balance not increased")
	}
}

func TestCharge_account(t *testing.T) {
	err := Charge_account("1", 1)
	if err != nil || Accounts["1"].balance != 0 {
		t.Log(err)
		t.Error("Balance not decreased")
	}
	err = Charge_account("1", 10)
	if !errors.Is(ErrAcountBelowZ, err) {
		t.Error("Balance decreased below zero")
	}

}
