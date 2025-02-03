package main

import "fmt"

func main() {
	numbers := []int{2, 4, 6}
	doubled := multiplyNumbers(&numbers, double)
	fmt.Println(doubled)
	tripled := multiplyNumbers(&numbers, triple)
	fmt.Println(tripled)
}

func multiplyNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
