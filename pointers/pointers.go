package main

import "fmt"

func main() {
	age := 34 // regular variable

	agePointer := &age

	fmt.Println("Age:", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println("Adult Years:", age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
