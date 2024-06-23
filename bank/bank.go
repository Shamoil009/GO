package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}
	return balance, nil

}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {

	// var accountBalance float64 = 1000
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		panic("Can't continue, sorry")
	}

	fmt.Println("Welcome to Go Bank")
	// for i := 0; i < 2; i++ {
	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1: Check Balance")
		fmt.Println("2: Deposit Money")
		fmt.Println("3: Withdraw Money")
		fmt.Println("4: Exit")

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)

		} else {
			fmt.Println("Goodbye!")
			break
		}
	}
	fmt.Println("Thanks for choosing our bank")
}
