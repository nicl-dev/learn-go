package main

import (
	"fmt"
	"time"
)

// lower case for this type as its not used outside of this package
type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u user) outputUserDetails() {
	//...
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func newUser(firstName, lastName, birthdate string) user {
	return user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}

// we need the *user pointer here so the original user struct gets mutated and not a copy of it.
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {

	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := newUser(userFirstName, userLastName, userBirthdate)
	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
