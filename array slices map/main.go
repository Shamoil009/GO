package main

import "fmt"

func main() {
	var productNames [4]string
	productNames = [4]string{"A book"}
	prices := [4]float64{23.99, 32.4, 21.4, 54}
	fmt.Println(prices)
	productNames[2] = "A carpet"
	fmt.Println(productNames)

	// 1 index to 2
	featuredPrices := prices[1:3]

	// start index to 2
	// featuredPrices:=prices[:3]

	// start 1 index to end
	// featuredPrices:=prices[1:]

	fmt.Println(featuredPrices)
}
