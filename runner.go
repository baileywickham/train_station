package main

// cli runner
import (
	r "github.com/baileywickham/runner"
)

func print_all_users() {
	for _, account := range Get_all_accounts() {
		account.Print()
	}
}
func add_money_uuid(uuid, balance int) {
	account, err := account_by_uuid(uuid)
	if err != nil {
		println(err.Error())
	}
	account.Add_balance(balance)
}
func charge_account_uuid(uuid int) {
	account, err := account_by_uuid(uuid)
	if err != nil {
		println(err.Error())
	}
	err = account.Charge(-1)
	if err != nil {
		println(err.Error())
	}
}

func print_account_uuid(uuid int) {
	account, err := account_by_uuid(uuid)
	if err != nil {
		println(err.Error())
	}
	account.Print()
}

func runner() {
	shell := r.NewShell()
	c1 := r.Command{"create_account", create_account, "Create new account [name] [amount]"}
	c2 := r.Command{"add", add_money_uuid, "Add money to an account [uuid] [amount]"}
	c3 := r.Command{"use", charge_account_uuid, "Charge account $1 [uuid]"}
	c4 := r.Command{"print", print_account_uuid, "Print single account [uuid]"}
	c5 := r.Command{"print_all", print_all_users, "Print all accounts"}
	shell.Add_command(c1, c2, c3, c4, c5)
	shell.Start()
}
