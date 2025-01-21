package main

import (
	"errors"
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

// we need the *user pointer here so the original user struct gets mutated and not a copy of it.
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate are required.")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func main() {

	userFirstName, _ := getUserData("Please enter your first name: ")
	userLastName, _ := getUserData("Please enter your last name: ")
	userBirthdate, _ := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}
	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) (string, error) {
	fmt.Print(promptText)
	var value string
	_, err := fmt.Scanln(&value)
	if err != nil {
		return "", err
	}
	return value, nil
}
