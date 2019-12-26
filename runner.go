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
			account, err := Create_account(tokens[1], i)
			if err != nil {
				println(err.Error())
			}
			println("Account ID: ", account.uuid)
		case "a":
			// add money
			i, _ := strconv.Atoi(tokens[2])
			uuid, _ := strconv.Atoi(tokens[1])
			err := Add_balance(uuid, i)
			if err != nil {
				log.Println(err.Error())
			}
		case "u":
			// use transaction. Default 1 dollar
			uuid, _ := strconv.Atoi(tokens[1])
			err := Charge_account(uuid, 1)
			if err != nil {
				println(err.Error())
			}
		case "p":
			uuid, _ := strconv.Atoi(tokens[1])
			print_account(uuid)
		default:
			println("Help:")
			println("c : create account, name, amount")
			println("a : add money, uuid, amount")
			println("u : use card, uuid")
			println("p : print account, uuid")
		}
	}
}
