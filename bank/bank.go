package main

import (
	"fmt"
	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	// var accountBalance float64 = 1000
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		panic("Can't continue, sorry")
	}

	fmt.Println("Welcome to Go Bank")
	fmt.Println("Reach us 24/7 ", randomdata.PhoneNumber())
	// for i := 0; i < 2; i++ {
	for {

		presentOptions()

		var choice int
		fmt.Scan(&choice)

		// switch case in go
		// switch choice {
		// case 1:
		// case 2:
		// case 3:
		// case 4:
		// default:
		// }

		if choice == 1 {
			fmt.Println("Your account balance is: ", accountBalance)
		} else if choice == 2 {
			fmt.Println("Your deposit: ")
			var depositAmmount float64
			fmt.Scan(&depositAmmount)

			if depositAmmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance = accountBalance + depositAmmount
			fmt.Println("Balance updated. New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		} else if choice == 3 {
			fmt.Println("Widthdraw amount: ")
			var withdrawAmmount float64
			fmt.Scan(&withdrawAmmount)

			if withdrawAmmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmmount > accountBalance {
				fmt.Println("Your widthdraw amount is more than deposited amount")
				continue
			}
			accountBalance = accountBalance - withdrawAmmount
			fmt.Println("Amount withdrawed. Remaining balance:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		} else {
			fmt.Println("Goodbye!")
			break
		}
	}
	fmt.Println("Thanks for choosing our bank")
}
