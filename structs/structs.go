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
	adminEmail, err := getUserData("Please enter your email: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	adminPassword, err := getUserData("Please enter a password: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	// test outputs for a user
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	// additional test outputs for embedded struct Admin
	appAdmin, err := user.NewAdmin(adminEmail, adminPassword)
	if err != nil {
		fmt.Println(err)
		return
	}

	appAdmin.OutputUserDetails()
	appAdmin.ClearUserName()
	appAdmin.OutputUserDetails()
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
