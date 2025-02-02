package main

import "fmt"

func main() {
	userNames := make([]string, 2)
	fmt.Println(userNames)

	userNames[0] = "Julie"

	fmt.Println(userNames)
}
