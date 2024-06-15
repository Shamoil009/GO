package main

import "fmt"

func main() {

	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1: Check Balance")
	fmt.Println("2: Deposit Money")
	fmt.Println("3: Withdraw Money")
	fmt.Println("4: Exit")

	var choice int
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your account balance is: ", accountBalance)
	} else if choice == 2 {
		fmt.Println("Your deposit: ")
		var depositAmmount float64
		fmt.Scan(&depositAmmount)

		if depositAmmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")
			return
		}

		accountBalance = accountBalance + depositAmmount
		fmt.Println("Balance updated. New amount:", accountBalance)
	} else if choice == 3 {
		fmt.Println("Widthdraw amount: ")
		var withdrawAmmount float64
		fmt.Scan(&withdrawAmmount)

		if withdrawAmmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")
			return
		}

		if withdrawAmmount > accountBalance {
			fmt.Println("Your widthdraw amount is more than deposited amount")
			return
		}
		accountBalance = accountBalance - withdrawAmmount
		fmt.Println("Amount withdrawed. Remaining balance:", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}
