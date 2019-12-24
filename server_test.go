package main

import "testing"

func TestCreate_account(t testing.T) {
	account, err := Create_account("1", "test", 0)
	// only verify uuid cause im lazy
	if err != nil || account.uuid != "1" {
		t.Log(err)
		t.Fail("Account not created")
	}
}

func TestAdd_balance(t testing.T) {
	err := Add_balance("1", 10)
	if err != nil || Accounts["1"].balance != 10 {
		t.Log(err)
		t.Fail("Account balance not update")
	}
}
func main() {

}
