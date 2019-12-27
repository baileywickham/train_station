package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func runner() {
	reader := bufio.NewReader(os.Stdin)
	//	defer
	func() {
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
		switch tokens[0] {
		case "c":
			i, _ := strconv.Atoi(tokens[2])
			account := Create_account(tokens[1], i)
			println("Account ID: ", account.uuid)
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
			account.Charge_account(1)
		case "p":
			uuid, _ := strconv.Atoi(tokens[1])
			account, err := account_by_uuid(uuid)
			if err != nil {
				log.Println(err.Error())
			}
			account.Print_account()

		default:
			println("Help:")
			println("c : create account, name, amount")
			println("a : add money, uuid, amount")
			println("u : use card, uuid")
			println("p : print account, uuid")
		}
	}
}
