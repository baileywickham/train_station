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
	println("Entering shell")
	for {
		print(":|:  ")
		text, _ := reader.ReadString('\n')
		//text = strings.Replace(text, "\n", "", -1)
		tokens := strings.Fields(text)
		switch tokens[0] {
		case "c":
			i, _ := strconv.Atoi(tokens[3])
			_, err := Create_account(tokens[1], tokens[2], i)
			if err != nil {
				println(err.Error())
			}
		case "a":
			// add money
			i, _ := strconv.Atoi(tokens[2])
			err := Add_balance(tokens[1], i)
			if err != nil {
				log.Println(err.Error())
			}
		case "u":
			// use transaction. Default 1 dollar
			err := Charge_account(tokens[1], 1)
			if err != nil {
				println(err.Error())
			}
		case "p":
			print_account(tokens[1])
		default:
			println("Help:")
			println("c : create account, uuid, name, amount")
			println("a : add money, uuid, amount")
			println("u : use card, uuid")
			println("p : print account, uuid")
		}
	}
}
