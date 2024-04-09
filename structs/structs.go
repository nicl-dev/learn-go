package main

import (
	"fmt"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	user := user{firstName, lastName, birthdate}

	fmt.Println("User data:", user)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
