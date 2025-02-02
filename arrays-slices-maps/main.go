package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)
	fmt.Println(userNames)

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)
}
