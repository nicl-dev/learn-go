package main

import "fmt"

func main() {
	age := 34 // regular variable

	agePointer := &age

	fmt.Println("Age:", *agePointer)

	adultYears := getAdultYears(agePointer)
	fmt.Println("Adult Years:", adultYears)
}

func getAdultYears(age *int) int {
	return *age - 18
}
