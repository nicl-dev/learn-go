package main

import "fmt"

func main() {
	age := 34 // regular variable

	agePointer := &age

	fmt.Println("Age:", *agePointer)

	getAdultYears(agePointer)
	fmt.Println("Adult Years:", age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
