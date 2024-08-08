package main

import (
	"fmt"

	"example.com/struct/user"
)

func main() {
	userFirstName := getUserData("please enter your first name: ")
	userLastName := getUserData("please enter your last name: ")
	userBirthDate := getUserData("please enter your birth date (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthDate)

	// appUser = &user.User{
	// 	FirstName: "Shamoil",
	// }

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("admin@example.com", "test1234")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}

func getUserData(promtText string) string {
	fmt.Print(promtText)
	var value string
	fmt.Scanln(&value) //Scanln will stop taking input when press enter
	return value
}
