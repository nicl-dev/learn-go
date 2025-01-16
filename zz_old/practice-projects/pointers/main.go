package main

import "fmt"

func main() {
	age := 32 // regular variable

	agePointer := &age // pointer variable
	fmt.Println("Age:", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println("Adult years:", age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
