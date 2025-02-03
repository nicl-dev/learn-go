package main

import "fmt"

func main() {
	numbers := []int{5, 6, 7, 8}
	sum := sumup(1, 2, 3, 4)
	anotherSum := sumup(numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(numbers ...int) int {
	sum := 0

	for _, i := range numbers {
		sum += i
	}

	return sum
}
