package main

import "fmt"

func main() {
	age := 32
	var agePointer *int
	agePointer = &age // & to get address

	fmt.Println("Age: ", *agePointer) // * to get value from pointer address
	editAgeToAdultYears(agePointer)
	var value int64
	fmt.Scan(&value)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
