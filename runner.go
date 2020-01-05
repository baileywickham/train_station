package main

// cli runner
import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func print_all_users() {
	for _, account := range Get_all_accounts() {
		account.Print()
	}
}

func runner() {
	reader := bufio.NewReader(os.Stdin)
	defer func() {
		// function to loop on panics
		if r := recover(); r != nil {
			log.Println("Panic: ", r)
			runner()
		}
	}()

	println("Entering shell")
	for {
		print(":|:  ")
		text, _ := reader.ReadString('\n')
		//text = strings.Replace(text, "\n", "", -1)
		tokens := strings.Fields(text)
		if len(tokens) == 0 {
			print_help()
			continue
		}
		switch tokens[0] {
		case "c":
			i, _ := strconv.Atoi(tokens[2])
			account := NewAccount(tokens[1], i)
			println("Account ID: ", account.UUID)
		case "a":
			// add money
			i, _ := strconv.Atoi(tokens[2])
			uuid, _ := strconv.Atoi(tokens[1])
			account, err := account_by_uuid(uuid)
			if err != nil {
				log.Println(err.Error())
			}
			account.Add_balance(i)
		case "u":
			// use transaction. Default 1 dollar
			uuid, _ := strconv.Atoi(tokens[1])
			account, err := account_by_uuid(uuid)
			if err != nil {
				log.Println(err.Error())
			}
			err = account.Charge(1)
			if err != nil {
				log.Println(err.Error())
			}
		case "p":
			uuid, _ := strconv.Atoi(tokens[1])
			account, err := account_by_uuid(uuid)
			if err != nil {
				log.Println(err.Error())
			}
			account.Print()
		case "pa":
			print_all_users()
		case "pad":
			for _, account := range get_all_accounts_db() {
				account.Print()
			}

		default:
			print_help()
		}
	}
}
func print_help() {
	println("Help:")
	println("c : create account, name, amount")
	println("a : add money, uuid, amount")
	println("u : use card, uuid")
	println("p : print account, uuid")
	println("pa: print all accounts")
}
