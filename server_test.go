package main

import "testing"

var account Account

func TestCreate_account(t *testing.T) {
	account = Create_account("test", 0)
	if account.name != "test" || account.balance != 0 {
		t.Error("Account not created")
	}
	a, err := account_by_uuid(account.uuid)
	if err != nil || a != account {
		t.Log("err", err)
		println("original:")
		account.Print_account()
		println("new:")
		a.Print_account()
		t.Error("database different than created")
	}
}

func TestAdd_balance(t *testing.T) {
	account.Add_balance(1)
	a, err := account_by_uuid(account.uuid)
	if err != nil || a.balance != 1 {
		t.Error("Balance not increased")
	}
}

func TestCharge_account(t *testing.T) {
	account.Charge_account(1)
	a, err := account_by_uuid(account.uuid)
	if err != nil || a.balance != 0 {
		t.Log(err)
		t.Error("Balance not decreased")
	}

}

func TestDelete_user(t *testing.T) {
	err := Delete_user(account.uuid)
	if err != nil {
		t.Log(err)
		t.Error("Errored on deleting user")
	}
	_, err = account_by_uuid(account.uuid)
	if err != ErrAccountNotFound {
		t.Log(err)
		t.Error("Failed to delete user")
	}

}
