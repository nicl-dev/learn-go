package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)

	secondResult := add("hello ", "sir")
	fmt.Println(secondResult)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
