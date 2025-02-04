package main

import "fmt"

func main() {
	// Task 1)
	hobbies := [3]string{"gaming", "coding", "diving"}
	fmt.Println(hobbies)

	// Task 2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// Task 3)
	sliceOne := hobbies[0:2]
	fmt.Println(sliceOne)

	sliceTwo := hobbies[:2]
	fmt.Println(sliceTwo)

	// Task 4)
	sliceTwo = sliceTwo[1:3]
	fmt.Println(sliceTwo)

	// Task 5)
	goals := []string{"learn go", "create an ebook library"}
	fmt.Println(goals)

	// Task 6)
	goals[1] = "improve my go skills"
	goals = append(goals, "create a nice app")
	fmt.Println(goals)

	// Task 7))
	type Product struct {
		title string
		id    int
		price float64
	}

	productList := []Product{
		{
			title: "apple",
			id:    0,
			price: 0.99,
		},
		{
			title: "banana",
			id:    1,
			price: 1.99,
		},
	}
	fmt.Println(productList)
	productList = append(productList, Product{
		title: "cherry",
		id:    2,
		price: 0.10,
	})
	fmt.Println(productList)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
