package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {

	userFirstName, err := getUserData("Please enter your first name: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	userLastName, err := getUserData("Please enter your last name: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	userBirthdate, err := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	if err != nil {
		fmt.Println(err)
		return
	}

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}
	// ... do something awesome with that gathered data!

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
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
