package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Profit: ")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax rate: ")

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	// if err1 != nil || err2 != nil || err3 != nil {
	// 	fmt.Println("ERROR")
	// 	fmt.Println(err1)
	// 	return
	// }

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)

	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - (taxRate / 100))
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New(infoText + " value is less than or equal to 0")
	}
	return userInput, nil
}
