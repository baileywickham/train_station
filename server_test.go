package main

import "testing"

var account Account

func TestCreate_account(t *testing.T) {
	account = Create_account("test", 0)
	if account.Name != "test" || account.Balance != 0 {
		t.Error("Account not created")
	}
	a, err := account_by_uuid(account.UUID)
	if err != nil || a != account {
		t.Log("err", err)
		println("original:")
		account.Print()
		println("new:")
		a.Print()
		t.Error("database different than created")
	}
}

func TestAdd_balance(t *testing.T) {
	account.Add_balance(1)
	a, err := account_by_uuid(account.UUID)
	if err != nil || a.Balance != 1 {
		t.Error("Balance not increased")
	}
}

func TestCharge_account(t *testing.T) {
	account.Charge(1)
	a, err := account_by_uuid(account.UUID)
	if err != nil || a.Balance != 0 {
		t.Log(err)
		t.Error("Balance not decreased")
	}

}

func TestDelete_user(t *testing.T) {
	err := Delete_user(account.UUID)
	if err != nil {
		t.Log(err)
		t.Error("Errored on deleting user")
	}
	_, err = account_by_uuid(account.UUID)
	if err != ErrAccountNotFound {
		t.Log(err)
		t.Error("Failed to delete user")
	}

}
