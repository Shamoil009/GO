package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("please enter your first name: ")
	userLastName := getUserData("please enter your last name: ")
	userBirthDate := getUserData("please enter your birth date (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}
	outputUserDetails(&appUser)

}

func outputUserDetails(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)

}
func getUserData(promtText string) string {
	fmt.Print(promtText)
	var value string
	fmt.Scan(&value)
	return value
}
