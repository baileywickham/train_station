package main

import "testing"

var account Account

func TestCreate_account(t *testing.T) {
	account = Create_account("test", 0)
	if account.name != "test" || account.balance != 0 {
		t.Error("Account not created")
	}
	a, err := account_by_uuid(account.uuid)
	if err != nil || a.uuid == account.uuid || a.name != account.name || a.balance != account.balance {
		t.Log(err)
		t.Error("database different than created")
	}
}

func TestAdd_balance(t *testing.T) {
	account.Add_balance(1)
	t.Error("Balance not increased")
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
