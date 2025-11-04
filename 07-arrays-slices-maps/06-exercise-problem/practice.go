package main

import "fmt"

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	hobbies := [3]string{"music", "reading", "traveling"}
	// 		Output (print) that array in the command line.
	fmt.Println("My hobbies:", hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	fmt.Println("First hobby:", hobbies[0])
	//		- The second and third element combined as a new list
	newList := []string{hobbies[1], hobbies[2]}
	fmt.Println("Other hobbies:", newList)

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	slice := hobbies[0:2]
	fmt.Println("1st and 2nd hobbies:", slice)
	//		Create that slice in two different ways (i.e. create two slices in the end)
	slice = hobbies[:2]
	fmt.Println("1st and 2nd hobbies:", slice)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	slice = hobbies[1:]
	fmt.Println("1st and 2nd hobbies:", slice)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	myGoals := []string{"Learning Go", "REST"}
	fmt.Println("My course goals:", myGoals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	myGoals[1] = "REST API"
	myGoals = append(myGoals, "Go on backend")
	fmt.Println("My course goals:", myGoals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.

	products := []Product{{1, "Prod1", 10}, {2, "Prod2", 20}}
	products = append(products, Product{3, "Prod3", 30})
	fmt.Println("My products:", products)
}

type Product struct {
	id    int
	title string
	price float64
}
