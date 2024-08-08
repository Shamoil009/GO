package user

import (
	"errors"
	"fmt"
	"time"
)

// to export values like firstName it need to be capitilize
// like this FirstName
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

// you can make a func belong to specific struct by turning it to method
// it will be available in struct also
func (u User) OutputUserDetails() {
	// func outputUserDetails(u *User) {

	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)

}

// using pointer to mutate original data
// without pointer it will mutate copy and not original data
func (u *User) ClearUserName() {

	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

func NewUser(firstName, lastName, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birth date are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil

}
