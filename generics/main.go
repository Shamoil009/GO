package main

import "fmt"

func main() {
	result := add(3, 4)
	fmt.Println("Result: ", result)
	result2 := add("Nice ","One")
	fmt.Println("Result: ", result2)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
