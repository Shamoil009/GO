package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var inverstmentAmount float64
	var years float64
	expectedReturnRate := 5.5 // to avoid var use :=

	fmt.Print("Investment value: ")
	fmt.Scan(&inverstmentAmount)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(inverstmentAmount, expectedReturnRate, years)

	// fmt.Println("Future Value: ", futureValue)
	// fmt.Println("Future Value (adjusted for inflation)", futureRealValue)
	// fmt.Printf("Future Value: %v\nFuture Value (adjusted for inflation): %v", futureValue, futureRealValue)
	fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for inflation): %.1f", futureValue, futureRealValue)

}

func calculateFutureValues(inverstmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := inverstmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
