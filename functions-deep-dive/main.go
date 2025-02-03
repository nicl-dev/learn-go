package main

import "fmt"

func main() {
	sum := sumup(1, 2, 3, 4)

	fmt.Println(sum)
}

func sumup(numbers ...int) int {
	sum := 0

	for _, i := range numbers {
		sum += i
	}

	return sum
}
