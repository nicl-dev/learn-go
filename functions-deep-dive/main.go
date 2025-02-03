package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 6, 7, 8}
	doubled := multiplyNumbers(&numbers, double)
	fmt.Println(doubled)
	tripled := multiplyNumbers(&numbers, triple)
	fmt.Println(tripled)

	transformFn1 := getTransfomerFunction(&numbers)
	transformFn2 := getTransfomerFunction(&moreNumbers)

	transformedNumbers := multiplyNumbers(&numbers, transformFn1)
	moreTransformedNumbers := multiplyNumbers(&moreNumbers, transformFn2)
	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

func multiplyNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTransfomerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
