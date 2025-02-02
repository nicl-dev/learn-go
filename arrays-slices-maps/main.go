package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)
	fmt.Println(userNames)

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRatings := make(map[string]float64, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7
	courseRatings["vue"] = 4.9
	fmt.Println(courseRatings)
}
