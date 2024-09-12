package main

import "fmt"

// array with dynamic length
func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99)
	fmt.Println(prices)

}

func main1() {
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
	// featuredPrices[0]=199.99 // this will assign value to parent array (prices)

	fmt.Println(featuredPrices)
}
